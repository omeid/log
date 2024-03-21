package log

import (
	"fmt"
	"log"
	"os"

	"github.com/fatih/color"
)

type Log interface {
	Log(v ...interface{})
	Logf(format string, v ...interface{})

	Info(v ...interface{})
	Infof(format string, v ...interface{})

	Notice(v ...interface{})
	Noticef(format string, v ...interface{})

	Warn(v ...interface{})
	Warnf(format string, v ...interface{})

	Error(v ...interface{})
	Errorf(format string, v ...interface{})

	Fatal(v ...interface{})
	Fatalf(format string, v ...interface{})

	New(string) Log
}

var Flags int

func New(prefix string) Log {
	l := log.New(os.Stdout, prefix, Flags)
	return &logger{l, ""}
}

type printer interface {
	Printf(string, ...interface{})
}

type logger struct {
	printer printer
	prefix  string
}

func (l *logger) New(prefix string) Log {
	return &logger{l.printer, l.prefix + prefix}
}

var bold = color.New(color.Bold).SprintfFunc()

func (l *logger) Log(v ...interface{}) {
	l.printer.Printf("%s%s ", l.prefix, fmt.Sprint(v...))
}

func (l *logger) Logf(format string, v ...interface{}) {
	l.Log(fmt.Sprintf(format, v...))
}

func (l *logger) Notice(v ...interface{}) {
	l.printer.Printf(bold("[NOTE] %s%s ", l.prefix, fmt.Sprint(v...)))
}

func (l *logger) Noticef(format string, v ...interface{}) {
	l.Notice(fmt.Sprintf(format, v...))
}

func (l *logger) Info(v ...interface{}) {
	l.printer.Printf("[INFO] %s%s ", l.prefix, fmt.Sprint(v...))
}

func (l *logger) Infof(format string, v ...interface{}) {
	l.Info(fmt.Sprintf(format, v...))
}

func (l *logger) Warn(v ...interface{}) {
	l.printer.Printf(color.YellowString("[WARN] %s%s ", l.prefix, fmt.Sprint(v...)))
}

func (l *logger) Warnf(format string, v ...interface{}) {
	l.Warn(fmt.Sprintf(format, v...))
}

func (l *logger) Error(v ...interface{}) {
	l.printer.Printf(color.RedString("[ERR!] %s%s ", l.prefix, fmt.Sprint(v...)))
}

func (l *logger) Errorf(format string, v ...interface{}) {
	l.Error(fmt.Sprintf(format, v...))
}

func (l *logger) Fatal(v ...interface{}) {
	l.printer.Printf(color.RedString("[FATAL] %s%s ", l.prefix, fmt.Sprint(v...)))
	os.Exit(1)
}

func (l *logger) Fatalf(format string, v ...interface{}) {
	l.Fatal(fmt.Sprintf(format, v...))
}
