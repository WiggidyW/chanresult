When using the recv-multi functions, it's highly recommended that the types themselves be pointers if the type is any bigger than an 'int' in memory.

If this is unfeasible, it's recommended to use interfaces and reflection instead.
