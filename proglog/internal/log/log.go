package log

import (
	"fmt"
	"io/ioutil"
	"os"
	"path"
	api "proglog/api/v1"
	"sort"
	"strconv"
	"strings"
	"sync"
)

func NewLog(dir string, c Config) (*Log, error) {
	if c.Segment.MaxStoreBytes == 0 {
		c.Segment.MaxStoreBytes = 1024
	}

	if c.Segment.MaxIndexBytes == 0 {
		c.Segment.MaxIndexBytes = 1024
	}

	l := &Log{
		Dir:    dir,
		Config: c,
	}
	return l, l.setup()
}

type Log struct {
	mu            sync.RWMutex
	Dir           string
	Config        Config
	activeSegment *segment
	segments      []*segment
}

func (l *Log) setup() error {
	files, err := ioutil.ReadDir(l.Dir)
	if err != nil {
		return err
	}

	var baseOffsets []uint64
	for _, file := range files {
		offStr := strings.TrimSuffix(
			file.Name(),
			path.Ext(file.Name()),
		)
		off, _ := strconv.ParseUint(offStr, 10, 0)
		baseOffsets = append(baseOffsets, off)
	}

	sort.Slice(baseOffsets, func(i, j int) bool {
		return baseOffsets[i] < baseOffsets[j]
	})

	for i := 0; i < len(baseOffsets); i++ {
		err = l.newSegment(baseOffsets[i])
		if err != nil {
			return err
		}
		// baseOffsets contains dup for index and store so we skip the dup
		i++
	}

	if l.segments == nil {
		err = l.newSegment(
			l.Config.Segment.InitialOffset,
		)
		if err != nil {
			return err
		}
	}

	return nil
}

func (l *Log) Append(record *api.Record) (uint64, error) {
	l.mu.Lock()
	defer l.mu.Unlock()

	off, err := l.activeSegment.Append(record)
	if err != nil {
		return 0, err
	}

	if l.activeSegment.IsMaxed() {
		err = l.newSegment(off + 1)
	}

	return off, err
}

func (l *Log) Read(off uint64) (*api.Record, error) {
	l.mu.RLock()
	defer l.mu.RUnlock()

	var s *segment
	for _, segment := range l.segments {
		if segment.baseOffset <= off && off < segment.nextOffset {
			s = segment
			break
		}
	}

	if s == nil || s.nextOffset <= off {
		return nil, fmt.Errorf("offset out of range: %d", off)
	}

	return s.Read(off)
}

func (l *Log) Close() error {
	l.mu.Lock()
	defer l.mu.Unlock()

	for _, segment := range l.segments {
		err := segment.Close()
		if err != nil {
			return err
		}
	}

	return nil
}

func (l *Log) Remove() error {
	err := l.Close()
	if err != nil {
		return err
	}
	return os.RemoveAll(l.Dir)
}

func (l *Log) Reset() error {
	err := l.Remove()
	if err != nil {
		return err
	}
	return l.setup()
}
