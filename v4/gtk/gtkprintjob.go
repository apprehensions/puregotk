// Package gtk was automatically generated by github.com/jwijenbergh/puregotk DO NOT EDIT
package gtk

import (
	"github.com/jwijenbergh/purego"
	"github.com/jwijenbergh/puregotk/internal/core"
	"github.com/jwijenbergh/puregotk/v4/cairo"
	"github.com/jwijenbergh/puregotk/v4/glib"
	"github.com/jwijenbergh/puregotk/v4/gobject"
)

// The type of callback that is passed to gtk_print_job_send().
//
// It is called when the print job has been completely sent.
type PrintJobCompleteFunc func(uintptr, uintptr, *glib.Error)

// A `GtkPrintJob` object represents a job that is sent to a printer.
//
// You only need to deal directly with print jobs if you use the
// non-portable [class@Gtk.PrintUnixDialog] API.
//
// Use [method@Gtk.PrintJob.get_surface] to obtain the cairo surface
// onto which the pages must be drawn. Use [method@Gtk.PrintJob.send]
// to send the finished job to the printer. If you don’t use cairo
// `GtkPrintJob` also supports printing of manually generated PostScript,
// via [method@Gtk.PrintJob.set_source_file].
type PrintJob struct {
	gobject.Object
}

func PrintJobNewFromInternalPtr(ptr uintptr) *PrintJob {
	cls := &PrintJob{}
	cls.Ptr = ptr
	return cls
}

var xNewPrintJob func(string, uintptr, uintptr, uintptr) uintptr

// Creates a new `GtkPrintJob`.
func NewPrintJob(TitleVar string, PrinterVar *Printer, SettingsVar *PrintSettings, PageSetupVar *PageSetup) *PrintJob {
	NewPrintJobPtr := xNewPrintJob(TitleVar, PrinterVar.GoPointer(), SettingsVar.GoPointer(), PageSetupVar.GoPointer())
	if NewPrintJobPtr == 0 {
		return nil
	}

	NewPrintJobCls := &PrintJob{}
	NewPrintJobCls.Ptr = NewPrintJobPtr
	return NewPrintJobCls
}

var xPrintJobGetCollate func(uintptr) bool

// Gets whether this job is printed collated.
func (x *PrintJob) GetCollate() bool {

	return xPrintJobGetCollate(x.GoPointer())

}

var xPrintJobGetNUp func(uintptr) uint

// Gets the n-up setting for this job.
func (x *PrintJob) GetNUp() uint {

	return xPrintJobGetNUp(x.GoPointer())

}

var xPrintJobGetNUpLayout func(uintptr) NumberUpLayout

// Gets the n-up layout setting for this job.
func (x *PrintJob) GetNUpLayout() NumberUpLayout {

	return xPrintJobGetNUpLayout(x.GoPointer())

}

var xPrintJobGetNumCopies func(uintptr) int32

// Gets the number of copies of this job.
func (x *PrintJob) GetNumCopies() int32 {

	return xPrintJobGetNumCopies(x.GoPointer())

}

var xPrintJobGetPageRanges func(uintptr, int32) uintptr

// Gets the page ranges for this job.
func (x *PrintJob) GetPageRanges(NRangesVar int32) uintptr {

	return xPrintJobGetPageRanges(x.GoPointer(), NRangesVar)

}

var xPrintJobGetPageSet func(uintptr) PageSet

// Gets the `GtkPageSet` setting for this job.
func (x *PrintJob) GetPageSet() PageSet {

	return xPrintJobGetPageSet(x.GoPointer())

}

var xPrintJobGetPages func(uintptr) PrintPages

// Gets the `GtkPrintPages` setting for this job.
func (x *PrintJob) GetPages() PrintPages {

	return xPrintJobGetPages(x.GoPointer())

}

var xPrintJobGetPrinter func(uintptr) uintptr

// Gets the `GtkPrinter` of the print job.
func (x *PrintJob) GetPrinter() *Printer {

	GetPrinterPtr := xPrintJobGetPrinter(x.GoPointer())
	if GetPrinterPtr == 0 {
		return nil
	}

	gobject.IncreaseRef(GetPrinterPtr)

	GetPrinterCls := &Printer{}
	GetPrinterCls.Ptr = GetPrinterPtr
	return GetPrinterCls

}

var xPrintJobGetReverse func(uintptr) bool

// Gets whether this job is printed reversed.
func (x *PrintJob) GetReverse() bool {

	return xPrintJobGetReverse(x.GoPointer())

}

var xPrintJobGetRotate func(uintptr) bool

// Gets whether the job is printed rotated.
func (x *PrintJob) GetRotate() bool {

	return xPrintJobGetRotate(x.GoPointer())

}

var xPrintJobGetScale func(uintptr) float64

// Gets the scale for this job.
func (x *PrintJob) GetScale() float64 {

	return xPrintJobGetScale(x.GoPointer())

}

var xPrintJobGetSettings func(uintptr) uintptr

// Gets the `GtkPrintSettings` of the print job.
func (x *PrintJob) GetSettings() *PrintSettings {

	GetSettingsPtr := xPrintJobGetSettings(x.GoPointer())
	if GetSettingsPtr == 0 {
		return nil
	}

	gobject.IncreaseRef(GetSettingsPtr)

	GetSettingsCls := &PrintSettings{}
	GetSettingsCls.Ptr = GetSettingsPtr
	return GetSettingsCls

}

var xPrintJobGetStatus func(uintptr) PrintStatus

// Gets the status of the print job.
func (x *PrintJob) GetStatus() PrintStatus {

	return xPrintJobGetStatus(x.GoPointer())

}

var xPrintJobGetSurface func(uintptr) *cairo.Surface

// Gets a cairo surface onto which the pages of
// the print job should be rendered.
func (x *PrintJob) GetSurface() *cairo.Surface {

	return xPrintJobGetSurface(x.GoPointer())

}

var xPrintJobGetTitle func(uintptr) string

// Gets the job title.
func (x *PrintJob) GetTitle() string {

	return xPrintJobGetTitle(x.GoPointer())

}

var xPrintJobGetTrackPrintStatus func(uintptr) bool

// Returns whether jobs will be tracked after printing.
//
// For details, see [method@Gtk.PrintJob.set_track_print_status].
func (x *PrintJob) GetTrackPrintStatus() bool {

	return xPrintJobGetTrackPrintStatus(x.GoPointer())

}

var xPrintJobSend func(uintptr, uintptr, uintptr, uintptr)

// Sends the print job off to the printer.
func (x *PrintJob) Send(CallbackVar PrintJobCompleteFunc, UserDataVar uintptr, DnotifyVar glib.DestroyNotify) {

	xPrintJobSend(x.GoPointer(), purego.NewCallback(CallbackVar), UserDataVar, purego.NewCallback(DnotifyVar))

}

var xPrintJobSetCollate func(uintptr, bool)

// Sets whether this job is printed collated.
func (x *PrintJob) SetCollate(CollateVar bool) {

	xPrintJobSetCollate(x.GoPointer(), CollateVar)

}

var xPrintJobSetNUp func(uintptr, uint)

// Sets the n-up setting for this job.
func (x *PrintJob) SetNUp(NUpVar uint) {

	xPrintJobSetNUp(x.GoPointer(), NUpVar)

}

var xPrintJobSetNUpLayout func(uintptr, NumberUpLayout)

// Sets the n-up layout setting for this job.
func (x *PrintJob) SetNUpLayout(LayoutVar NumberUpLayout) {

	xPrintJobSetNUpLayout(x.GoPointer(), LayoutVar)

}

var xPrintJobSetNumCopies func(uintptr, int32)

// Sets the number of copies for this job.
func (x *PrintJob) SetNumCopies(NumCopiesVar int32) {

	xPrintJobSetNumCopies(x.GoPointer(), NumCopiesVar)

}

var xPrintJobSetPageRanges func(uintptr, uintptr, int32)

// Sets the page ranges for this job.
func (x *PrintJob) SetPageRanges(RangesVar uintptr, NRangesVar int32) {

	xPrintJobSetPageRanges(x.GoPointer(), RangesVar, NRangesVar)

}

var xPrintJobSetPageSet func(uintptr, PageSet)

// Sets the `GtkPageSet` setting for this job.
func (x *PrintJob) SetPageSet(PageSetVar PageSet) {

	xPrintJobSetPageSet(x.GoPointer(), PageSetVar)

}

var xPrintJobSetPages func(uintptr, PrintPages)

// Sets the `GtkPrintPages` setting for this job.
func (x *PrintJob) SetPages(PagesVar PrintPages) {

	xPrintJobSetPages(x.GoPointer(), PagesVar)

}

var xPrintJobSetReverse func(uintptr, bool)

// Sets whether this job is printed reversed.
func (x *PrintJob) SetReverse(ReverseVar bool) {

	xPrintJobSetReverse(x.GoPointer(), ReverseVar)

}

var xPrintJobSetRotate func(uintptr, bool)

// Sets whether this job is printed rotated.
func (x *PrintJob) SetRotate(RotateVar bool) {

	xPrintJobSetRotate(x.GoPointer(), RotateVar)

}

var xPrintJobSetScale func(uintptr, float64)

// Sets the scale for this job.
//
// 1.0 means unscaled.
func (x *PrintJob) SetScale(ScaleVar float64) {

	xPrintJobSetScale(x.GoPointer(), ScaleVar)

}

var xPrintJobSetSourceFd func(uintptr, int32) bool

// Make the `GtkPrintJob` send an existing document to the
// printing system.
//
// The file can be in any format understood by the platforms
// printing system (typically PostScript, but on many platforms
// PDF may work too). See [method@Gtk.Printer.accepts_pdf] and
// [method@Gtk.Printer.accepts_ps].
//
// This is similar to [method@Gtk.PrintJob.set_source_file],
// but takes expects an open file descriptor for the file,
// instead of a filename.
func (x *PrintJob) SetSourceFd(FdVar int32) bool {

	return xPrintJobSetSourceFd(x.GoPointer(), FdVar)

}

var xPrintJobSetSourceFile func(uintptr, string) bool

// Make the `GtkPrintJob` send an existing document to the
// printing system.
//
// The file can be in any format understood by the platforms
// printing system (typically PostScript, but on many platforms
// PDF may work too). See [method@Gtk.Printer.accepts_pdf] and
// [method@Gtk.Printer.accepts_ps].
func (x *PrintJob) SetSourceFile(FilenameVar string) bool {

	return xPrintJobSetSourceFile(x.GoPointer(), FilenameVar)

}

var xPrintJobSetTrackPrintStatus func(uintptr, bool)

// If track_status is %TRUE, the print job will try to continue report
// on the status of the print job in the printer queues and printer.
//
// This can allow your application to show things like “out of paper”
// issues, and when the print job actually reaches the printer.
//
// This function is often implemented using some form of polling,
// so it should not be enabled unless needed.
func (x *PrintJob) SetTrackPrintStatus(TrackStatusVar bool) {

	xPrintJobSetTrackPrintStatus(x.GoPointer(), TrackStatusVar)

}

func (c *PrintJob) GoPointer() uintptr {
	return c.Ptr
}

func (c *PrintJob) SetGoPointer(ptr uintptr) {
	c.Ptr = ptr
}

// Emitted when the status of a job changes.
//
// The signal handler can use [method@Gtk.PrintJob.get_status]
// to obtain the new status.
func (x *PrintJob) ConnectStatusChanged(cb func(PrintJob)) {
	fcb := func(clsPtr uintptr) {
		fa := PrintJob{}
		fa.Ptr = clsPtr

		cb(fa)

	}
	gobject.ObjectConnect(x.GoPointer(), "signal::status-changed", purego.NewCallback(fcb))
}

func init() {
	lib, err := purego.Dlopen(core.GetPath("GTK"), purego.RTLD_NOW|purego.RTLD_GLOBAL)
	if err != nil {
		panic(err)
	}

	core.PuregoSafeRegister(&xNewPrintJob, lib, "gtk_print_job_new")

	core.PuregoSafeRegister(&xPrintJobGetCollate, lib, "gtk_print_job_get_collate")
	core.PuregoSafeRegister(&xPrintJobGetNUp, lib, "gtk_print_job_get_n_up")
	core.PuregoSafeRegister(&xPrintJobGetNUpLayout, lib, "gtk_print_job_get_n_up_layout")
	core.PuregoSafeRegister(&xPrintJobGetNumCopies, lib, "gtk_print_job_get_num_copies")
	core.PuregoSafeRegister(&xPrintJobGetPageRanges, lib, "gtk_print_job_get_page_ranges")
	core.PuregoSafeRegister(&xPrintJobGetPageSet, lib, "gtk_print_job_get_page_set")
	core.PuregoSafeRegister(&xPrintJobGetPages, lib, "gtk_print_job_get_pages")
	core.PuregoSafeRegister(&xPrintJobGetPrinter, lib, "gtk_print_job_get_printer")
	core.PuregoSafeRegister(&xPrintJobGetReverse, lib, "gtk_print_job_get_reverse")
	core.PuregoSafeRegister(&xPrintJobGetRotate, lib, "gtk_print_job_get_rotate")
	core.PuregoSafeRegister(&xPrintJobGetScale, lib, "gtk_print_job_get_scale")
	core.PuregoSafeRegister(&xPrintJobGetSettings, lib, "gtk_print_job_get_settings")
	core.PuregoSafeRegister(&xPrintJobGetStatus, lib, "gtk_print_job_get_status")
	core.PuregoSafeRegister(&xPrintJobGetSurface, lib, "gtk_print_job_get_surface")
	core.PuregoSafeRegister(&xPrintJobGetTitle, lib, "gtk_print_job_get_title")
	core.PuregoSafeRegister(&xPrintJobGetTrackPrintStatus, lib, "gtk_print_job_get_track_print_status")
	core.PuregoSafeRegister(&xPrintJobSend, lib, "gtk_print_job_send")
	core.PuregoSafeRegister(&xPrintJobSetCollate, lib, "gtk_print_job_set_collate")
	core.PuregoSafeRegister(&xPrintJobSetNUp, lib, "gtk_print_job_set_n_up")
	core.PuregoSafeRegister(&xPrintJobSetNUpLayout, lib, "gtk_print_job_set_n_up_layout")
	core.PuregoSafeRegister(&xPrintJobSetNumCopies, lib, "gtk_print_job_set_num_copies")
	core.PuregoSafeRegister(&xPrintJobSetPageRanges, lib, "gtk_print_job_set_page_ranges")
	core.PuregoSafeRegister(&xPrintJobSetPageSet, lib, "gtk_print_job_set_page_set")
	core.PuregoSafeRegister(&xPrintJobSetPages, lib, "gtk_print_job_set_pages")
	core.PuregoSafeRegister(&xPrintJobSetReverse, lib, "gtk_print_job_set_reverse")
	core.PuregoSafeRegister(&xPrintJobSetRotate, lib, "gtk_print_job_set_rotate")
	core.PuregoSafeRegister(&xPrintJobSetScale, lib, "gtk_print_job_set_scale")
	core.PuregoSafeRegister(&xPrintJobSetSourceFd, lib, "gtk_print_job_set_source_fd")
	core.PuregoSafeRegister(&xPrintJobSetSourceFile, lib, "gtk_print_job_set_source_file")
	core.PuregoSafeRegister(&xPrintJobSetTrackPrintStatus, lib, "gtk_print_job_set_track_print_status")

}
