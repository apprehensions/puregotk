// Package gtk was automatically generated by github.com/jwijenbergh/puregotk DO NOT EDIT
package gtk

import (
	"github.com/jwijenbergh/purego"
	"github.com/jwijenbergh/puregotk/internal/core"
	"github.com/jwijenbergh/puregotk/v4/glib"
	"github.com/jwijenbergh/puregotk/v4/gobject"
)

// The type of function passed to gtk_enumerate_printers().
//
// Note that you need to ref @printer, if you want to keep
// a reference to it after the function has returned.
type PrinterFunc func(uintptr, uintptr) bool

type PrintBackend struct {
}

// Specifies which features the print dialog should offer.
//
// If neither %GTK_PRINT_CAPABILITY_GENERATE_PDF nor
// %GTK_PRINT_CAPABILITY_GENERATE_PS is specified, GTK assumes that all
// formats are supported.
type PrintCapabilities int

const (

	// Print dialog will offer printing even/odd pages.
	PrintCapabilityPageSetValue PrintCapabilities = 1
	// Print dialog will allow to print multiple copies.
	PrintCapabilityCopiesValue PrintCapabilities = 2
	// Print dialog will allow to collate multiple copies.
	PrintCapabilityCollateValue PrintCapabilities = 4
	// Print dialog will allow to print pages in reverse order.
	PrintCapabilityReverseValue PrintCapabilities = 8
	// Print dialog will allow to scale the output.
	PrintCapabilityScaleValue PrintCapabilities = 16
	// The program will send the document to
	//   the printer in PDF format
	PrintCapabilityGeneratePdfValue PrintCapabilities = 32
	// The program will send the document to
	//   the printer in Postscript format
	PrintCapabilityGeneratePsValue PrintCapabilities = 64
	// Print dialog will offer a preview
	PrintCapabilityPreviewValue PrintCapabilities = 128
	// Print dialog will offer printing multiple
	//   pages per sheet
	PrintCapabilityNumberUpValue PrintCapabilities = 256
	// Print dialog will allow to rearrange
	//   pages when printing multiple pages per sheet
	PrintCapabilityNumberUpLayoutValue PrintCapabilities = 512
)

var xEnumeratePrinters func(uintptr, uintptr, uintptr, bool)

// Calls a function for all `GtkPrinter`s.
//
// If @func returns %TRUE, the enumeration is stopped.
func EnumeratePrinters(FuncVar PrinterFunc, DataVar uintptr, DestroyVar glib.DestroyNotify, WaitVar bool) {

	xEnumeratePrinters(purego.NewCallback(FuncVar), DataVar, purego.NewCallback(DestroyVar), WaitVar)

}

// A `GtkPrinter` object represents a printer.
//
// You only need to deal directly with printers if you use the
// non-portable [class@Gtk.PrintUnixDialog] API.
//
// A `GtkPrinter` allows to get status information about the printer,
// such as its description, its location, the number of queued jobs,
// etc. Most importantly, a `GtkPrinter` object can be used to create
// a [class@Gtk.PrintJob] object, which lets you print to the printer.
type Printer struct {
	gobject.Object
}

func PrinterNewFromInternalPtr(ptr uintptr) *Printer {
	cls := &Printer{}
	cls.Ptr = ptr
	return cls
}

var xNewPrinter func(string, *PrintBackend, bool) uintptr

// Creates a new `GtkPrinter`.
func NewPrinter(NameVar string, BackendVar *PrintBackend, VirtualVar bool) *Printer {
	var cls *Printer

	cret := xNewPrinter(NameVar, BackendVar, VirtualVar)

	if cret == 0 {
		return cls
	}
	cls = &Printer{}
	cls.Ptr = cret
	return cls
}

var xPrinterAcceptsPdf func(uintptr) bool

// Returns whether the printer accepts input in
// PDF format.
func (x *Printer) AcceptsPdf() bool {

	cret := xPrinterAcceptsPdf(x.GoPointer())
	return cret
}

var xPrinterAcceptsPs func(uintptr) bool

// Returns whether the printer accepts input in
// PostScript format.
func (x *Printer) AcceptsPs() bool {

	cret := xPrinterAcceptsPs(x.GoPointer())
	return cret
}

var xPrinterCompare func(uintptr, uintptr) int

// Compares two printers.
func (x *Printer) Compare(BVar *Printer) int {

	cret := xPrinterCompare(x.GoPointer(), BVar.GoPointer())
	return cret
}

var xPrinterGetBackend func(uintptr) *PrintBackend

// Returns the backend of the printer.
func (x *Printer) GetBackend() *PrintBackend {

	cret := xPrinterGetBackend(x.GoPointer())
	return cret
}

var xPrinterGetCapabilities func(uintptr) PrintCapabilities

// Returns the printer’s capabilities.
//
// This is useful when you’re using `GtkPrintUnixDialog`’s
// manual-capabilities setting and need to know which settings
// the printer can handle and which you must handle yourself.
//
// This will return 0 unless the printer’s details are
// available, see [method@Gtk.Printer.has_details] and
// [method@Gtk.Printer.request_details].
func (x *Printer) GetCapabilities() PrintCapabilities {

	cret := xPrinterGetCapabilities(x.GoPointer())
	return cret
}

var xPrinterGetDefaultPageSize func(uintptr) uintptr

// Returns default page size of @printer.
func (x *Printer) GetDefaultPageSize() *PageSetup {
	var cls *PageSetup

	cret := xPrinterGetDefaultPageSize(x.GoPointer())

	if cret == 0 {
		return cls
	}
	cls = &PageSetup{}
	cls.Ptr = cret
	return cls
}

var xPrinterGetDescription func(uintptr) string

// Gets the description of the printer.
func (x *Printer) GetDescription() string {

	cret := xPrinterGetDescription(x.GoPointer())
	return cret
}

var xPrinterGetHardMargins func(uintptr, float64, float64, float64, float64) bool

// Retrieve the hard margins of @printer.
//
// These are the margins that define the area at the borders
// of the paper that the printer cannot print to.
//
// Note: This will not succeed unless the printer’s details are
// available, see [method@Gtk.Printer.has_details] and
// [method@Gtk.Printer.request_details].
func (x *Printer) GetHardMargins(TopVar float64, BottomVar float64, LeftVar float64, RightVar float64) bool {

	cret := xPrinterGetHardMargins(x.GoPointer(), TopVar, BottomVar, LeftVar, RightVar)
	return cret
}

var xPrinterGetHardMarginsForPaperSize func(uintptr, *PaperSize, float64, float64, float64, float64) bool

// Retrieve the hard margins of @printer for @paper_size.
//
// These are the margins that define the area at the borders
// of the paper that the printer cannot print to.
//
// Note: This will not succeed unless the printer’s details are
// available, see [method@Gtk.Printer.has_details] and
// [method@Gtk.Printer.request_details].
func (x *Printer) GetHardMarginsForPaperSize(PaperSizeVar *PaperSize, TopVar float64, BottomVar float64, LeftVar float64, RightVar float64) bool {

	cret := xPrinterGetHardMarginsForPaperSize(x.GoPointer(), PaperSizeVar, TopVar, BottomVar, LeftVar, RightVar)
	return cret
}

var xPrinterGetIconName func(uintptr) string

// Gets the name of the icon to use for the printer.
func (x *Printer) GetIconName() string {

	cret := xPrinterGetIconName(x.GoPointer())
	return cret
}

var xPrinterGetJobCount func(uintptr) int

// Gets the number of jobs currently queued on the printer.
func (x *Printer) GetJobCount() int {

	cret := xPrinterGetJobCount(x.GoPointer())
	return cret
}

var xPrinterGetLocation func(uintptr) string

// Returns a description of the location of the printer.
func (x *Printer) GetLocation() string {

	cret := xPrinterGetLocation(x.GoPointer())
	return cret
}

var xPrinterGetName func(uintptr) string

// Returns the name of the printer.
func (x *Printer) GetName() string {

	cret := xPrinterGetName(x.GoPointer())
	return cret
}

var xPrinterGetStateMessage func(uintptr) string

// Returns the state message describing the current state
// of the printer.
func (x *Printer) GetStateMessage() string {

	cret := xPrinterGetStateMessage(x.GoPointer())
	return cret
}

var xPrinterHasDetails func(uintptr) bool

// Returns whether the printer details are available.
func (x *Printer) HasDetails() bool {

	cret := xPrinterHasDetails(x.GoPointer())
	return cret
}

var xPrinterIsAcceptingJobs func(uintptr) bool

// Returns whether the printer is accepting jobs
func (x *Printer) IsAcceptingJobs() bool {

	cret := xPrinterIsAcceptingJobs(x.GoPointer())
	return cret
}

var xPrinterIsActive func(uintptr) bool

// Returns whether the printer is currently active (i.e.
// accepts new jobs).
func (x *Printer) IsActive() bool {

	cret := xPrinterIsActive(x.GoPointer())
	return cret
}

var xPrinterIsDefault func(uintptr) bool

// Returns whether the printer is the default printer.
func (x *Printer) IsDefault() bool {

	cret := xPrinterIsDefault(x.GoPointer())
	return cret
}

var xPrinterIsPaused func(uintptr) bool

// Returns whether the printer is currently paused.
//
// A paused printer still accepts jobs, but it is not
// printing them.
func (x *Printer) IsPaused() bool {

	cret := xPrinterIsPaused(x.GoPointer())
	return cret
}

var xPrinterIsVirtual func(uintptr) bool

// Returns whether the printer is virtual (i.e. does not
// represent actual printer hardware, but something like
// a CUPS class).
func (x *Printer) IsVirtual() bool {

	cret := xPrinterIsVirtual(x.GoPointer())
	return cret
}

var xPrinterListPapers func(uintptr) *glib.List

// Lists all the paper sizes @printer supports.
//
// This will return and empty list unless the printer’s details
// are available, see [method@Gtk.Printer.has_details] and
// [method@Gtk.Printer.request_details].
func (x *Printer) ListPapers() *glib.List {

	cret := xPrinterListPapers(x.GoPointer())
	return cret
}

var xPrinterRequestDetails func(uintptr)

// Requests the printer details.
//
// When the details are available, the
// [signal@Gtk.Printer::details-acquired] signal
// will be emitted on @printer.
func (x *Printer) RequestDetails() {

	xPrinterRequestDetails(x.GoPointer())

}

func (c *Printer) GoPointer() uintptr {
	return c.Ptr
}

func (c *Printer) SetGoPointer(ptr uintptr) {
	c.Ptr = ptr
}

// Emitted in response to a request for detailed information
// about a printer from the print backend.
//
// The @success parameter indicates if the information was
// actually obtained.
func (x *Printer) ConnectDetailsAcquired(cb func(Printer, bool)) {
	fcb := func(clsPtr uintptr, SuccessVarp bool) {
		fa := Printer{}
		fa.Ptr = clsPtr

		cb(fa, SuccessVarp)

	}
	gobject.ObjectConnect(x.GoPointer(), "signal::details-acquired", purego.NewCallback(fcb))
}

func init() {
	lib, err := purego.Dlopen(core.GetPath("GTK"), purego.RTLD_NOW|purego.RTLD_GLOBAL)
	if err != nil {
		panic(err)
	}
	core.PuregoSafeRegister(&xEnumeratePrinters, lib, "gtk_enumerate_printers")

	core.PuregoSafeRegister(&xNewPrinter, lib, "gtk_printer_new")

	core.PuregoSafeRegister(&xPrinterAcceptsPdf, lib, "gtk_printer_accepts_pdf")
	core.PuregoSafeRegister(&xPrinterAcceptsPs, lib, "gtk_printer_accepts_ps")
	core.PuregoSafeRegister(&xPrinterCompare, lib, "gtk_printer_compare")
	core.PuregoSafeRegister(&xPrinterGetBackend, lib, "gtk_printer_get_backend")
	core.PuregoSafeRegister(&xPrinterGetCapabilities, lib, "gtk_printer_get_capabilities")
	core.PuregoSafeRegister(&xPrinterGetDefaultPageSize, lib, "gtk_printer_get_default_page_size")
	core.PuregoSafeRegister(&xPrinterGetDescription, lib, "gtk_printer_get_description")
	core.PuregoSafeRegister(&xPrinterGetHardMargins, lib, "gtk_printer_get_hard_margins")
	core.PuregoSafeRegister(&xPrinterGetHardMarginsForPaperSize, lib, "gtk_printer_get_hard_margins_for_paper_size")
	core.PuregoSafeRegister(&xPrinterGetIconName, lib, "gtk_printer_get_icon_name")
	core.PuregoSafeRegister(&xPrinterGetJobCount, lib, "gtk_printer_get_job_count")
	core.PuregoSafeRegister(&xPrinterGetLocation, lib, "gtk_printer_get_location")
	core.PuregoSafeRegister(&xPrinterGetName, lib, "gtk_printer_get_name")
	core.PuregoSafeRegister(&xPrinterGetStateMessage, lib, "gtk_printer_get_state_message")
	core.PuregoSafeRegister(&xPrinterHasDetails, lib, "gtk_printer_has_details")
	core.PuregoSafeRegister(&xPrinterIsAcceptingJobs, lib, "gtk_printer_is_accepting_jobs")
	core.PuregoSafeRegister(&xPrinterIsActive, lib, "gtk_printer_is_active")
	core.PuregoSafeRegister(&xPrinterIsDefault, lib, "gtk_printer_is_default")
	core.PuregoSafeRegister(&xPrinterIsPaused, lib, "gtk_printer_is_paused")
	core.PuregoSafeRegister(&xPrinterIsVirtual, lib, "gtk_printer_is_virtual")
	core.PuregoSafeRegister(&xPrinterListPapers, lib, "gtk_printer_list_papers")
	core.PuregoSafeRegister(&xPrinterRequestDetails, lib, "gtk_printer_request_details")

}
