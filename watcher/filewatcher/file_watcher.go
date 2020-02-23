package filewatcher

import (
	"log"
	"time"

	"github.com/radovskyb/watcher"
)

//ProcessFunc 监听事件的处理函数
type ProcessFunc func(*watcher.Watcher)

type FileWatcher struct {
	fileWatcher *watcher.Watcher
	processFunc ProcessFunc
}

func (w *FileWatcher) runFileWatcher(fileName string, processer ProcessFunc, interval time.Duration) {
	w.fileWatcher = watcher.New()
	w.fileWatcher.SetMaxEvents(1)
	w.fileWatcher.FilterOps(watcher.Write)
	w.fileWatcher.Add(fileName)
	w.processFunc = processer
	go w.processFunc(w.fileWatcher)
	if err := w.fileWatcher.Start(interval); err != nil {
		log.Fatalln(err)
	}
}

//RunFileWatcher 启动文件监控
func RunFileWatcher(fileName string, processer ProcessFunc, interval time.Duration) *FileWatcher {
	w := &FileWatcher{}
	go w.runFileWatcher(fileName, processer, interval)
	return w
}
