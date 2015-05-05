package raft

import (
	"log"
	"github.com/hashicorp/govector/govec"
)

type WrapperLogger struct {
	logger *log.Logger
	vec_logger *govec.GoLog
}

func (w *WrapperLogger) print(msg string) {
	w.logger.Printf(msg)
	w.vec_logger.LogLocalEvent(msg)
}

func (w *WrapperLogger) PrepareSend(msg string, payload []byte) []byte {
	return w.vec_logger.PrepareSend(msg, payload)
}

func (w *WrapperLogger) UnpackReceive(msg string, payload []byte) {
	w.vec_logger.UnpackReceive(msg, payload)
}

func (w *WrapperLogger) DisableLogging() {
	w.vec_logger.DisableLogging();
}
