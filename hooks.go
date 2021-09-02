package ipset

// HookFunc is a function suitable to be set as the GlobalHook, and will be
// called at the beginning of each ipset operation.
//
// For example, this would run a timer metric duration on each set function:
//
//     FIXME
type HookFunc func(set IPSet, action string) HookCleanupFunc

// HookCleanupFunc is a function returned by a HookFunc, deferred until the
// end of the ipset operation.  The error, if any, of the function being
// hooked is made available as a parameter, to allow any post-processing to
// take into account the success (or not) of the set operation.
type HookCleanupFunc func(error)

var GlobalHook HookFunc
