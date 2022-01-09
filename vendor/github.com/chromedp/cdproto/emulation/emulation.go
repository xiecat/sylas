// Package emulation provides the Chrome DevTools Protocol
// commands, types, and events for the Emulation domain.
//
// This domain emulates different environments for the page.
//
// Generated by the cdproto-gen command.
package emulation

// Code generated by cdproto-gen. DO NOT EDIT.

import (
	"context"

	"github.com/chromedp/cdproto/cdp"
	"github.com/chromedp/cdproto/page"
)

// CanEmulateParams tells whether emulation is supported.
type CanEmulateParams struct{}

// CanEmulate tells whether emulation is supported.
//
// See: https://chromedevtools.github.io/devtools-protocol/tot/Emulation#method-canEmulate
func CanEmulate() *CanEmulateParams {
	return &CanEmulateParams{}
}

// CanEmulateReturns return values.
type CanEmulateReturns struct {
	Result bool `json:"result,omitempty"` // True if emulation is supported.
}

// Do executes Emulation.canEmulate against the provided context.
//
// returns:
//   result - True if emulation is supported.
func (p *CanEmulateParams) Do(ctx context.Context) (result bool, err error) {
	// execute
	var res CanEmulateReturns
	err = cdp.Execute(ctx, CommandCanEmulate, nil, &res)
	if err != nil {
		return false, err
	}

	return res.Result, nil
}

// ClearDeviceMetricsOverrideParams clears the overridden device metrics.
type ClearDeviceMetricsOverrideParams struct{}

// ClearDeviceMetricsOverride clears the overridden device metrics.
//
// See: https://chromedevtools.github.io/devtools-protocol/tot/Emulation#method-clearDeviceMetricsOverride
func ClearDeviceMetricsOverride() *ClearDeviceMetricsOverrideParams {
	return &ClearDeviceMetricsOverrideParams{}
}

// Do executes Emulation.clearDeviceMetricsOverride against the provided context.
func (p *ClearDeviceMetricsOverrideParams) Do(ctx context.Context) (err error) {
	return cdp.Execute(ctx, CommandClearDeviceMetricsOverride, nil, nil)
}

// ClearGeolocationOverrideParams clears the overridden Geolocation Position
// and Error.
type ClearGeolocationOverrideParams struct{}

// ClearGeolocationOverride clears the overridden Geolocation Position and
// Error.
//
// See: https://chromedevtools.github.io/devtools-protocol/tot/Emulation#method-clearGeolocationOverride
func ClearGeolocationOverride() *ClearGeolocationOverrideParams {
	return &ClearGeolocationOverrideParams{}
}

// Do executes Emulation.clearGeolocationOverride against the provided context.
func (p *ClearGeolocationOverrideParams) Do(ctx context.Context) (err error) {
	return cdp.Execute(ctx, CommandClearGeolocationOverride, nil, nil)
}

// ResetPageScaleFactorParams requests that page scale factor is reset to
// initial values.
type ResetPageScaleFactorParams struct{}

// ResetPageScaleFactor requests that page scale factor is reset to initial
// values.
//
// See: https://chromedevtools.github.io/devtools-protocol/tot/Emulation#method-resetPageScaleFactor
func ResetPageScaleFactor() *ResetPageScaleFactorParams {
	return &ResetPageScaleFactorParams{}
}

// Do executes Emulation.resetPageScaleFactor against the provided context.
func (p *ResetPageScaleFactorParams) Do(ctx context.Context) (err error) {
	return cdp.Execute(ctx, CommandResetPageScaleFactor, nil, nil)
}

// SetFocusEmulationEnabledParams enables or disables simulating a focused
// and active page.
type SetFocusEmulationEnabledParams struct {
	Enabled bool `json:"enabled"` // Whether to enable to disable focus emulation.
}

// SetFocusEmulationEnabled enables or disables simulating a focused and
// active page.
//
// See: https://chromedevtools.github.io/devtools-protocol/tot/Emulation#method-setFocusEmulationEnabled
//
// parameters:
//   enabled - Whether to enable to disable focus emulation.
func SetFocusEmulationEnabled(enabled bool) *SetFocusEmulationEnabledParams {
	return &SetFocusEmulationEnabledParams{
		Enabled: enabled,
	}
}

// Do executes Emulation.setFocusEmulationEnabled against the provided context.
func (p *SetFocusEmulationEnabledParams) Do(ctx context.Context) (err error) {
	return cdp.Execute(ctx, CommandSetFocusEmulationEnabled, p, nil)
}

// SetAutoDarkModeOverrideParams automatically render all web contents using
// a dark theme.
type SetAutoDarkModeOverrideParams struct {
	Enabled bool `json:"enabled,omitempty"` // Whether to enable or disable automatic dark mode. If not specified, any existing override will be cleared.
}

// SetAutoDarkModeOverride automatically render all web contents using a dark
// theme.
//
// See: https://chromedevtools.github.io/devtools-protocol/tot/Emulation#method-setAutoDarkModeOverride
//
// parameters:
func SetAutoDarkModeOverride() *SetAutoDarkModeOverrideParams {
	return &SetAutoDarkModeOverrideParams{}
}

// WithEnabled whether to enable or disable automatic dark mode. If not
// specified, any existing override will be cleared.
func (p SetAutoDarkModeOverrideParams) WithEnabled(enabled bool) *SetAutoDarkModeOverrideParams {
	p.Enabled = enabled
	return &p
}

// Do executes Emulation.setAutoDarkModeOverride against the provided context.
func (p *SetAutoDarkModeOverrideParams) Do(ctx context.Context) (err error) {
	return cdp.Execute(ctx, CommandSetAutoDarkModeOverride, p, nil)
}

// SetCPUThrottlingRateParams enables CPU throttling to emulate slow CPUs.
type SetCPUThrottlingRateParams struct {
	Rate float64 `json:"rate"` // Throttling rate as a slowdown factor (1 is no throttle, 2 is 2x slowdown, etc).
}

// SetCPUThrottlingRate enables CPU throttling to emulate slow CPUs.
//
// See: https://chromedevtools.github.io/devtools-protocol/tot/Emulation#method-setCPUThrottlingRate
//
// parameters:
//   rate - Throttling rate as a slowdown factor (1 is no throttle, 2 is 2x slowdown, etc).
func SetCPUThrottlingRate(rate float64) *SetCPUThrottlingRateParams {
	return &SetCPUThrottlingRateParams{
		Rate: rate,
	}
}

// Do executes Emulation.setCPUThrottlingRate against the provided context.
func (p *SetCPUThrottlingRateParams) Do(ctx context.Context) (err error) {
	return cdp.Execute(ctx, CommandSetCPUThrottlingRate, p, nil)
}

// SetDefaultBackgroundColorOverrideParams sets or clears an override of the
// default background color of the frame. This override is used if the content
// does not specify one.
type SetDefaultBackgroundColorOverrideParams struct {
	Color *cdp.RGBA `json:"color,omitempty"` // RGBA of the default background color. If not specified, any existing override will be cleared.
}

// SetDefaultBackgroundColorOverride sets or clears an override of the
// default background color of the frame. This override is used if the content
// does not specify one.
//
// See: https://chromedevtools.github.io/devtools-protocol/tot/Emulation#method-setDefaultBackgroundColorOverride
//
// parameters:
func SetDefaultBackgroundColorOverride() *SetDefaultBackgroundColorOverrideParams {
	return &SetDefaultBackgroundColorOverrideParams{}
}

// WithColor rGBA of the default background color. If not specified, any
// existing override will be cleared.
func (p SetDefaultBackgroundColorOverrideParams) WithColor(color *cdp.RGBA) *SetDefaultBackgroundColorOverrideParams {
	p.Color = color
	return &p
}

// Do executes Emulation.setDefaultBackgroundColorOverride against the provided context.
func (p *SetDefaultBackgroundColorOverrideParams) Do(ctx context.Context) (err error) {
	return cdp.Execute(ctx, CommandSetDefaultBackgroundColorOverride, p, nil)
}

// SetDeviceMetricsOverrideParams overrides the values of device screen
// dimensions (window.screen.width, window.screen.height, window.innerWidth,
// window.innerHeight, and "device-width"/"device-height"-related CSS media
// query results).
type SetDeviceMetricsOverrideParams struct {
	Width              int64              `json:"width"`                        // Overriding width value in pixels (minimum 0, maximum 10000000). 0 disables the override.
	Height             int64              `json:"height"`                       // Overriding height value in pixels (minimum 0, maximum 10000000). 0 disables the override.
	DeviceScaleFactor  float64            `json:"deviceScaleFactor"`            // Overriding device scale factor value. 0 disables the override.
	Mobile             bool               `json:"mobile"`                       // Whether to emulate mobile device. This includes viewport meta tag, overlay scrollbars, text autosizing and more.
	Scale              float64            `json:"scale,omitempty"`              // Scale to apply to resulting view image.
	ScreenWidth        int64              `json:"screenWidth,omitempty"`        // Overriding screen width value in pixels (minimum 0, maximum 10000000).
	ScreenHeight       int64              `json:"screenHeight,omitempty"`       // Overriding screen height value in pixels (minimum 0, maximum 10000000).
	PositionX          int64              `json:"positionX,omitempty"`          // Overriding view X position on screen in pixels (minimum 0, maximum 10000000).
	PositionY          int64              `json:"positionY,omitempty"`          // Overriding view Y position on screen in pixels (minimum 0, maximum 10000000).
	DontSetVisibleSize bool               `json:"dontSetVisibleSize,omitempty"` // Do not set visible view size, rely upon explicit setVisibleSize call.
	ScreenOrientation  *ScreenOrientation `json:"screenOrientation,omitempty"`  // Screen orientation override.
	Viewport           *page.Viewport     `json:"viewport,omitempty"`           // If set, the visible area of the page will be overridden to this viewport. This viewport change is not observed by the page, e.g. viewport-relative elements do not change positions.
	DisplayFeature     *DisplayFeature    `json:"displayFeature,omitempty"`     // If set, the display feature of a multi-segment screen. If not set, multi-segment support is turned-off.
}

// SetDeviceMetricsOverride overrides the values of device screen dimensions
// (window.screen.width, window.screen.height, window.innerWidth,
// window.innerHeight, and "device-width"/"device-height"-related CSS media
// query results).
//
// See: https://chromedevtools.github.io/devtools-protocol/tot/Emulation#method-setDeviceMetricsOverride
//
// parameters:
//   width - Overriding width value in pixels (minimum 0, maximum 10000000). 0 disables the override.
//   height - Overriding height value in pixels (minimum 0, maximum 10000000). 0 disables the override.
//   deviceScaleFactor - Overriding device scale factor value. 0 disables the override.
//   mobile - Whether to emulate mobile device. This includes viewport meta tag, overlay scrollbars, text autosizing and more.
func SetDeviceMetricsOverride(width int64, height int64, deviceScaleFactor float64, mobile bool) *SetDeviceMetricsOverrideParams {
	return &SetDeviceMetricsOverrideParams{
		Width:             width,
		Height:            height,
		DeviceScaleFactor: deviceScaleFactor,
		Mobile:            mobile,
	}
}

// WithScale scale to apply to resulting view image.
func (p SetDeviceMetricsOverrideParams) WithScale(scale float64) *SetDeviceMetricsOverrideParams {
	p.Scale = scale
	return &p
}

// WithScreenWidth overriding screen width value in pixels (minimum 0,
// maximum 10000000).
func (p SetDeviceMetricsOverrideParams) WithScreenWidth(screenWidth int64) *SetDeviceMetricsOverrideParams {
	p.ScreenWidth = screenWidth
	return &p
}

// WithScreenHeight overriding screen height value in pixels (minimum 0,
// maximum 10000000).
func (p SetDeviceMetricsOverrideParams) WithScreenHeight(screenHeight int64) *SetDeviceMetricsOverrideParams {
	p.ScreenHeight = screenHeight
	return &p
}

// WithPositionX overriding view X position on screen in pixels (minimum 0,
// maximum 10000000).
func (p SetDeviceMetricsOverrideParams) WithPositionX(positionX int64) *SetDeviceMetricsOverrideParams {
	p.PositionX = positionX
	return &p
}

// WithPositionY overriding view Y position on screen in pixels (minimum 0,
// maximum 10000000).
func (p SetDeviceMetricsOverrideParams) WithPositionY(positionY int64) *SetDeviceMetricsOverrideParams {
	p.PositionY = positionY
	return &p
}

// WithDontSetVisibleSize do not set visible view size, rely upon explicit
// setVisibleSize call.
func (p SetDeviceMetricsOverrideParams) WithDontSetVisibleSize(dontSetVisibleSize bool) *SetDeviceMetricsOverrideParams {
	p.DontSetVisibleSize = dontSetVisibleSize
	return &p
}

// WithScreenOrientation screen orientation override.
func (p SetDeviceMetricsOverrideParams) WithScreenOrientation(screenOrientation *ScreenOrientation) *SetDeviceMetricsOverrideParams {
	p.ScreenOrientation = screenOrientation
	return &p
}

// WithViewport if set, the visible area of the page will be overridden to
// this viewport. This viewport change is not observed by the page, e.g.
// viewport-relative elements do not change positions.
func (p SetDeviceMetricsOverrideParams) WithViewport(viewport *page.Viewport) *SetDeviceMetricsOverrideParams {
	p.Viewport = viewport
	return &p
}

// WithDisplayFeature if set, the display feature of a multi-segment screen.
// If not set, multi-segment support is turned-off.
func (p SetDeviceMetricsOverrideParams) WithDisplayFeature(displayFeature *DisplayFeature) *SetDeviceMetricsOverrideParams {
	p.DisplayFeature = displayFeature
	return &p
}

// Do executes Emulation.setDeviceMetricsOverride against the provided context.
func (p *SetDeviceMetricsOverrideParams) Do(ctx context.Context) (err error) {
	return cdp.Execute(ctx, CommandSetDeviceMetricsOverride, p, nil)
}

// SetScrollbarsHiddenParams [no description].
type SetScrollbarsHiddenParams struct {
	Hidden bool `json:"hidden"` // Whether scrollbars should be always hidden.
}

// SetScrollbarsHidden [no description].
//
// See: https://chromedevtools.github.io/devtools-protocol/tot/Emulation#method-setScrollbarsHidden
//
// parameters:
//   hidden - Whether scrollbars should be always hidden.
func SetScrollbarsHidden(hidden bool) *SetScrollbarsHiddenParams {
	return &SetScrollbarsHiddenParams{
		Hidden: hidden,
	}
}

// Do executes Emulation.setScrollbarsHidden against the provided context.
func (p *SetScrollbarsHiddenParams) Do(ctx context.Context) (err error) {
	return cdp.Execute(ctx, CommandSetScrollbarsHidden, p, nil)
}

// SetDocumentCookieDisabledParams [no description].
type SetDocumentCookieDisabledParams struct {
	Disabled bool `json:"disabled"` // Whether document.coookie API should be disabled.
}

// SetDocumentCookieDisabled [no description].
//
// See: https://chromedevtools.github.io/devtools-protocol/tot/Emulation#method-setDocumentCookieDisabled
//
// parameters:
//   disabled - Whether document.coookie API should be disabled.
func SetDocumentCookieDisabled(disabled bool) *SetDocumentCookieDisabledParams {
	return &SetDocumentCookieDisabledParams{
		Disabled: disabled,
	}
}

// Do executes Emulation.setDocumentCookieDisabled against the provided context.
func (p *SetDocumentCookieDisabledParams) Do(ctx context.Context) (err error) {
	return cdp.Execute(ctx, CommandSetDocumentCookieDisabled, p, nil)
}

// SetEmitTouchEventsForMouseParams [no description].
type SetEmitTouchEventsForMouseParams struct {
	Enabled       bool                                    `json:"enabled"`                 // Whether touch emulation based on mouse input should be enabled.
	Configuration SetEmitTouchEventsForMouseConfiguration `json:"configuration,omitempty"` // Touch/gesture events configuration. Default: current platform.
}

// SetEmitTouchEventsForMouse [no description].
//
// See: https://chromedevtools.github.io/devtools-protocol/tot/Emulation#method-setEmitTouchEventsForMouse
//
// parameters:
//   enabled - Whether touch emulation based on mouse input should be enabled.
func SetEmitTouchEventsForMouse(enabled bool) *SetEmitTouchEventsForMouseParams {
	return &SetEmitTouchEventsForMouseParams{
		Enabled: enabled,
	}
}

// WithConfiguration touch/gesture events configuration. Default: current
// platform.
func (p SetEmitTouchEventsForMouseParams) WithConfiguration(configuration SetEmitTouchEventsForMouseConfiguration) *SetEmitTouchEventsForMouseParams {
	p.Configuration = configuration
	return &p
}

// Do executes Emulation.setEmitTouchEventsForMouse against the provided context.
func (p *SetEmitTouchEventsForMouseParams) Do(ctx context.Context) (err error) {
	return cdp.Execute(ctx, CommandSetEmitTouchEventsForMouse, p, nil)
}

// SetEmulatedMediaParams emulates the given media type or media feature for
// CSS media queries.
type SetEmulatedMediaParams struct {
	Media    string          `json:"media,omitempty"`    // Media type to emulate. Empty string disables the override.
	Features []*MediaFeature `json:"features,omitempty"` // Media features to emulate.
}

// SetEmulatedMedia emulates the given media type or media feature for CSS
// media queries.
//
// See: https://chromedevtools.github.io/devtools-protocol/tot/Emulation#method-setEmulatedMedia
//
// parameters:
func SetEmulatedMedia() *SetEmulatedMediaParams {
	return &SetEmulatedMediaParams{}
}

// WithMedia media type to emulate. Empty string disables the override.
func (p SetEmulatedMediaParams) WithMedia(media string) *SetEmulatedMediaParams {
	p.Media = media
	return &p
}

// WithFeatures media features to emulate.
func (p SetEmulatedMediaParams) WithFeatures(features []*MediaFeature) *SetEmulatedMediaParams {
	p.Features = features
	return &p
}

// Do executes Emulation.setEmulatedMedia against the provided context.
func (p *SetEmulatedMediaParams) Do(ctx context.Context) (err error) {
	return cdp.Execute(ctx, CommandSetEmulatedMedia, p, nil)
}

// SetEmulatedVisionDeficiencyParams emulates the given vision deficiency.
type SetEmulatedVisionDeficiencyParams struct {
	Type SetEmulatedVisionDeficiencyType `json:"type"` // Vision deficiency to emulate.
}

// SetEmulatedVisionDeficiency emulates the given vision deficiency.
//
// See: https://chromedevtools.github.io/devtools-protocol/tot/Emulation#method-setEmulatedVisionDeficiency
//
// parameters:
//   type - Vision deficiency to emulate.
func SetEmulatedVisionDeficiency(typeVal SetEmulatedVisionDeficiencyType) *SetEmulatedVisionDeficiencyParams {
	return &SetEmulatedVisionDeficiencyParams{
		Type: typeVal,
	}
}

// Do executes Emulation.setEmulatedVisionDeficiency against the provided context.
func (p *SetEmulatedVisionDeficiencyParams) Do(ctx context.Context) (err error) {
	return cdp.Execute(ctx, CommandSetEmulatedVisionDeficiency, p, nil)
}

// SetGeolocationOverrideParams overrides the Geolocation Position or Error.
// Omitting any of the parameters emulates position unavailable.
type SetGeolocationOverrideParams struct {
	Latitude  float64 `json:"latitude,omitempty"`  // Mock latitude
	Longitude float64 `json:"longitude,omitempty"` // Mock longitude
	Accuracy  float64 `json:"accuracy,omitempty"`  // Mock accuracy
}

// SetGeolocationOverride overrides the Geolocation Position or Error.
// Omitting any of the parameters emulates position unavailable.
//
// See: https://chromedevtools.github.io/devtools-protocol/tot/Emulation#method-setGeolocationOverride
//
// parameters:
func SetGeolocationOverride() *SetGeolocationOverrideParams {
	return &SetGeolocationOverrideParams{}
}

// WithLatitude mock latitude.
func (p SetGeolocationOverrideParams) WithLatitude(latitude float64) *SetGeolocationOverrideParams {
	p.Latitude = latitude
	return &p
}

// WithLongitude mock longitude.
func (p SetGeolocationOverrideParams) WithLongitude(longitude float64) *SetGeolocationOverrideParams {
	p.Longitude = longitude
	return &p
}

// WithAccuracy mock accuracy.
func (p SetGeolocationOverrideParams) WithAccuracy(accuracy float64) *SetGeolocationOverrideParams {
	p.Accuracy = accuracy
	return &p
}

// Do executes Emulation.setGeolocationOverride against the provided context.
func (p *SetGeolocationOverrideParams) Do(ctx context.Context) (err error) {
	return cdp.Execute(ctx, CommandSetGeolocationOverride, p, nil)
}

// SetIdleOverrideParams overrides the Idle state.
type SetIdleOverrideParams struct {
	IsUserActive     bool `json:"isUserActive"`     // Mock isUserActive
	IsScreenUnlocked bool `json:"isScreenUnlocked"` // Mock isScreenUnlocked
}

// SetIdleOverride overrides the Idle state.
//
// See: https://chromedevtools.github.io/devtools-protocol/tot/Emulation#method-setIdleOverride
//
// parameters:
//   isUserActive - Mock isUserActive
//   isScreenUnlocked - Mock isScreenUnlocked
func SetIdleOverride(isUserActive bool, isScreenUnlocked bool) *SetIdleOverrideParams {
	return &SetIdleOverrideParams{
		IsUserActive:     isUserActive,
		IsScreenUnlocked: isScreenUnlocked,
	}
}

// Do executes Emulation.setIdleOverride against the provided context.
func (p *SetIdleOverrideParams) Do(ctx context.Context) (err error) {
	return cdp.Execute(ctx, CommandSetIdleOverride, p, nil)
}

// ClearIdleOverrideParams clears Idle state overrides.
type ClearIdleOverrideParams struct{}

// ClearIdleOverride clears Idle state overrides.
//
// See: https://chromedevtools.github.io/devtools-protocol/tot/Emulation#method-clearIdleOverride
func ClearIdleOverride() *ClearIdleOverrideParams {
	return &ClearIdleOverrideParams{}
}

// Do executes Emulation.clearIdleOverride against the provided context.
func (p *ClearIdleOverrideParams) Do(ctx context.Context) (err error) {
	return cdp.Execute(ctx, CommandClearIdleOverride, nil, nil)
}

// SetPageScaleFactorParams sets a specified page scale factor.
type SetPageScaleFactorParams struct {
	PageScaleFactor float64 `json:"pageScaleFactor"` // Page scale factor.
}

// SetPageScaleFactor sets a specified page scale factor.
//
// See: https://chromedevtools.github.io/devtools-protocol/tot/Emulation#method-setPageScaleFactor
//
// parameters:
//   pageScaleFactor - Page scale factor.
func SetPageScaleFactor(pageScaleFactor float64) *SetPageScaleFactorParams {
	return &SetPageScaleFactorParams{
		PageScaleFactor: pageScaleFactor,
	}
}

// Do executes Emulation.setPageScaleFactor against the provided context.
func (p *SetPageScaleFactorParams) Do(ctx context.Context) (err error) {
	return cdp.Execute(ctx, CommandSetPageScaleFactor, p, nil)
}

// SetScriptExecutionDisabledParams switches script execution in the page.
type SetScriptExecutionDisabledParams struct {
	Value bool `json:"value"` // Whether script execution should be disabled in the page.
}

// SetScriptExecutionDisabled switches script execution in the page.
//
// See: https://chromedevtools.github.io/devtools-protocol/tot/Emulation#method-setScriptExecutionDisabled
//
// parameters:
//   value - Whether script execution should be disabled in the page.
func SetScriptExecutionDisabled(value bool) *SetScriptExecutionDisabledParams {
	return &SetScriptExecutionDisabledParams{
		Value: value,
	}
}

// Do executes Emulation.setScriptExecutionDisabled against the provided context.
func (p *SetScriptExecutionDisabledParams) Do(ctx context.Context) (err error) {
	return cdp.Execute(ctx, CommandSetScriptExecutionDisabled, p, nil)
}

// SetTouchEmulationEnabledParams enables touch on platforms which do not
// support them.
type SetTouchEmulationEnabledParams struct {
	Enabled        bool  `json:"enabled"`                  // Whether the touch event emulation should be enabled.
	MaxTouchPoints int64 `json:"maxTouchPoints,omitempty"` // Maximum touch points supported. Defaults to one.
}

// SetTouchEmulationEnabled enables touch on platforms which do not support
// them.
//
// See: https://chromedevtools.github.io/devtools-protocol/tot/Emulation#method-setTouchEmulationEnabled
//
// parameters:
//   enabled - Whether the touch event emulation should be enabled.
func SetTouchEmulationEnabled(enabled bool) *SetTouchEmulationEnabledParams {
	return &SetTouchEmulationEnabledParams{
		Enabled: enabled,
	}
}

// WithMaxTouchPoints maximum touch points supported. Defaults to one.
func (p SetTouchEmulationEnabledParams) WithMaxTouchPoints(maxTouchPoints int64) *SetTouchEmulationEnabledParams {
	p.MaxTouchPoints = maxTouchPoints
	return &p
}

// Do executes Emulation.setTouchEmulationEnabled against the provided context.
func (p *SetTouchEmulationEnabledParams) Do(ctx context.Context) (err error) {
	return cdp.Execute(ctx, CommandSetTouchEmulationEnabled, p, nil)
}

// SetVirtualTimePolicyParams turns on virtual time for all frames (replacing
// real-time with a synthetic time source) and sets the current virtual time
// policy. Note this supersedes any previous time budget.
type SetVirtualTimePolicyParams struct {
	Policy                            VirtualTimePolicy   `json:"policy"`
	Budget                            float64             `json:"budget,omitempty"`                            // If set, after this many virtual milliseconds have elapsed virtual time will be paused and a virtualTimeBudgetExpired event is sent.
	MaxVirtualTimeTaskStarvationCount int64               `json:"maxVirtualTimeTaskStarvationCount,omitempty"` // If set this specifies the maximum number of tasks that can be run before virtual is forced forwards to prevent deadlock.
	WaitForNavigation                 bool                `json:"waitForNavigation,omitempty"`                 // If set the virtual time policy change should be deferred until any frame starts navigating. Note any previous deferred policy change is superseded.
	InitialVirtualTime                *cdp.TimeSinceEpoch `json:"initialVirtualTime,omitempty"`                // If set, base::Time::Now will be overridden to initially return this value.
}

// SetVirtualTimePolicy turns on virtual time for all frames (replacing
// real-time with a synthetic time source) and sets the current virtual time
// policy. Note this supersedes any previous time budget.
//
// See: https://chromedevtools.github.io/devtools-protocol/tot/Emulation#method-setVirtualTimePolicy
//
// parameters:
//   policy
func SetVirtualTimePolicy(policy VirtualTimePolicy) *SetVirtualTimePolicyParams {
	return &SetVirtualTimePolicyParams{
		Policy: policy,
	}
}

// WithBudget if set, after this many virtual milliseconds have elapsed
// virtual time will be paused and a virtualTimeBudgetExpired event is sent.
func (p SetVirtualTimePolicyParams) WithBudget(budget float64) *SetVirtualTimePolicyParams {
	p.Budget = budget
	return &p
}

// WithMaxVirtualTimeTaskStarvationCount if set this specifies the maximum
// number of tasks that can be run before virtual is forced forwards to prevent
// deadlock.
func (p SetVirtualTimePolicyParams) WithMaxVirtualTimeTaskStarvationCount(maxVirtualTimeTaskStarvationCount int64) *SetVirtualTimePolicyParams {
	p.MaxVirtualTimeTaskStarvationCount = maxVirtualTimeTaskStarvationCount
	return &p
}

// WithWaitForNavigation if set the virtual time policy change should be
// deferred until any frame starts navigating. Note any previous deferred policy
// change is superseded.
func (p SetVirtualTimePolicyParams) WithWaitForNavigation(waitForNavigation bool) *SetVirtualTimePolicyParams {
	p.WaitForNavigation = waitForNavigation
	return &p
}

// WithInitialVirtualTime if set, base::Time::Now will be overridden to
// initially return this value.
func (p SetVirtualTimePolicyParams) WithInitialVirtualTime(initialVirtualTime *cdp.TimeSinceEpoch) *SetVirtualTimePolicyParams {
	p.InitialVirtualTime = initialVirtualTime
	return &p
}

// SetVirtualTimePolicyReturns return values.
type SetVirtualTimePolicyReturns struct {
	VirtualTimeTicksBase float64 `json:"virtualTimeTicksBase,omitempty"` // Absolute timestamp at which virtual time was first enabled (up time in milliseconds).
}

// Do executes Emulation.setVirtualTimePolicy against the provided context.
//
// returns:
//   virtualTimeTicksBase - Absolute timestamp at which virtual time was first enabled (up time in milliseconds).
func (p *SetVirtualTimePolicyParams) Do(ctx context.Context) (virtualTimeTicksBase float64, err error) {
	// execute
	var res SetVirtualTimePolicyReturns
	err = cdp.Execute(ctx, CommandSetVirtualTimePolicy, p, &res)
	if err != nil {
		return 0, err
	}

	return res.VirtualTimeTicksBase, nil
}

// SetLocaleOverrideParams overrides default host system locale with the
// specified one.
type SetLocaleOverrideParams struct {
	Locale string `json:"locale,omitempty"` // ICU style C locale (e.g. "en_US"). If not specified or empty, disables the override and restores default host system locale.
}

// SetLocaleOverride overrides default host system locale with the specified
// one.
//
// See: https://chromedevtools.github.io/devtools-protocol/tot/Emulation#method-setLocaleOverride
//
// parameters:
func SetLocaleOverride() *SetLocaleOverrideParams {
	return &SetLocaleOverrideParams{}
}

// WithLocale iCU style C locale (e.g. "en_US"). If not specified or empty,
// disables the override and restores default host system locale.
func (p SetLocaleOverrideParams) WithLocale(locale string) *SetLocaleOverrideParams {
	p.Locale = locale
	return &p
}

// Do executes Emulation.setLocaleOverride against the provided context.
func (p *SetLocaleOverrideParams) Do(ctx context.Context) (err error) {
	return cdp.Execute(ctx, CommandSetLocaleOverride, p, nil)
}

// SetTimezoneOverrideParams overrides default host system timezone with the
// specified one.
type SetTimezoneOverrideParams struct {
	TimezoneID string `json:"timezoneId"` // The timezone identifier. If empty, disables the override and restores default host system timezone.
}

// SetTimezoneOverride overrides default host system timezone with the
// specified one.
//
// See: https://chromedevtools.github.io/devtools-protocol/tot/Emulation#method-setTimezoneOverride
//
// parameters:
//   timezoneID - The timezone identifier. If empty, disables the override and restores default host system timezone.
func SetTimezoneOverride(timezoneID string) *SetTimezoneOverrideParams {
	return &SetTimezoneOverrideParams{
		TimezoneID: timezoneID,
	}
}

// Do executes Emulation.setTimezoneOverride against the provided context.
func (p *SetTimezoneOverrideParams) Do(ctx context.Context) (err error) {
	return cdp.Execute(ctx, CommandSetTimezoneOverride, p, nil)
}

// SetDisabledImageTypesParams [no description].
type SetDisabledImageTypesParams struct {
	ImageTypes []DisabledImageType `json:"imageTypes"` // Image types to disable.
}

// SetDisabledImageTypes [no description].
//
// See: https://chromedevtools.github.io/devtools-protocol/tot/Emulation#method-setDisabledImageTypes
//
// parameters:
//   imageTypes - Image types to disable.
func SetDisabledImageTypes(imageTypes []DisabledImageType) *SetDisabledImageTypesParams {
	return &SetDisabledImageTypesParams{
		ImageTypes: imageTypes,
	}
}

// Do executes Emulation.setDisabledImageTypes against the provided context.
func (p *SetDisabledImageTypesParams) Do(ctx context.Context) (err error) {
	return cdp.Execute(ctx, CommandSetDisabledImageTypes, p, nil)
}

// SetUserAgentOverrideParams allows overriding user agent with the given
// string.
type SetUserAgentOverrideParams struct {
	UserAgent         string             `json:"userAgent"`                   // User agent to use.
	AcceptLanguage    string             `json:"acceptLanguage,omitempty"`    // Browser langugage to emulate.
	Platform          string             `json:"platform,omitempty"`          // The platform navigator.platform should return.
	UserAgentMetadata *UserAgentMetadata `json:"userAgentMetadata,omitempty"` // To be sent in Sec-CH-UA-* headers and returned in navigator.userAgentData
}

// SetUserAgentOverride allows overriding user agent with the given string.
//
// See: https://chromedevtools.github.io/devtools-protocol/tot/Emulation#method-setUserAgentOverride
//
// parameters:
//   userAgent - User agent to use.
func SetUserAgentOverride(userAgent string) *SetUserAgentOverrideParams {
	return &SetUserAgentOverrideParams{
		UserAgent: userAgent,
	}
}

// WithAcceptLanguage browser langugage to emulate.
func (p SetUserAgentOverrideParams) WithAcceptLanguage(acceptLanguage string) *SetUserAgentOverrideParams {
	p.AcceptLanguage = acceptLanguage
	return &p
}

// WithPlatform the platform navigator.platform should return.
func (p SetUserAgentOverrideParams) WithPlatform(platform string) *SetUserAgentOverrideParams {
	p.Platform = platform
	return &p
}

// WithUserAgentMetadata to be sent in Sec-CH-UA-* headers and returned in
// navigator.userAgentData.
func (p SetUserAgentOverrideParams) WithUserAgentMetadata(userAgentMetadata *UserAgentMetadata) *SetUserAgentOverrideParams {
	p.UserAgentMetadata = userAgentMetadata
	return &p
}

// Do executes Emulation.setUserAgentOverride against the provided context.
func (p *SetUserAgentOverrideParams) Do(ctx context.Context) (err error) {
	return cdp.Execute(ctx, CommandSetUserAgentOverride, p, nil)
}

// Command names.
const (
	CommandCanEmulate                        = "Emulation.canEmulate"
	CommandClearDeviceMetricsOverride        = "Emulation.clearDeviceMetricsOverride"
	CommandClearGeolocationOverride          = "Emulation.clearGeolocationOverride"
	CommandResetPageScaleFactor              = "Emulation.resetPageScaleFactor"
	CommandSetFocusEmulationEnabled          = "Emulation.setFocusEmulationEnabled"
	CommandSetAutoDarkModeOverride           = "Emulation.setAutoDarkModeOverride"
	CommandSetCPUThrottlingRate              = "Emulation.setCPUThrottlingRate"
	CommandSetDefaultBackgroundColorOverride = "Emulation.setDefaultBackgroundColorOverride"
	CommandSetDeviceMetricsOverride          = "Emulation.setDeviceMetricsOverride"
	CommandSetScrollbarsHidden               = "Emulation.setScrollbarsHidden"
	CommandSetDocumentCookieDisabled         = "Emulation.setDocumentCookieDisabled"
	CommandSetEmitTouchEventsForMouse        = "Emulation.setEmitTouchEventsForMouse"
	CommandSetEmulatedMedia                  = "Emulation.setEmulatedMedia"
	CommandSetEmulatedVisionDeficiency       = "Emulation.setEmulatedVisionDeficiency"
	CommandSetGeolocationOverride            = "Emulation.setGeolocationOverride"
	CommandSetIdleOverride                   = "Emulation.setIdleOverride"
	CommandClearIdleOverride                 = "Emulation.clearIdleOverride"
	CommandSetPageScaleFactor                = "Emulation.setPageScaleFactor"
	CommandSetScriptExecutionDisabled        = "Emulation.setScriptExecutionDisabled"
	CommandSetTouchEmulationEnabled          = "Emulation.setTouchEmulationEnabled"
	CommandSetVirtualTimePolicy              = "Emulation.setVirtualTimePolicy"
	CommandSetLocaleOverride                 = "Emulation.setLocaleOverride"
	CommandSetTimezoneOverride               = "Emulation.setTimezoneOverride"
	CommandSetDisabledImageTypes             = "Emulation.setDisabledImageTypes"
	CommandSetUserAgentOverride              = "Emulation.setUserAgentOverride"
)
