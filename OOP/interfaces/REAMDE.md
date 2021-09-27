INTERFACES
==================

link: https://jordanorelli.com/post/32665860244/how-to-use-interfaces-in-go

One of the core concepts in the Go's type system is that instead of designing our abstractions in terms of what kind of data our systems can hold, we design our abstractions in terms of what actions our types can execute.

The interface {} type, the empty interface, is the source of much confusion. Since there's no implements keyword, all types implement at least zero methos, and satisfying an interface is done automatically. All types satisfy the empty interface.