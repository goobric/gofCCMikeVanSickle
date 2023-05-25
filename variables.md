# Summary

## Variable Declaration
*   var foo int
*   var bar int = 42
*   bat := 89

*   Can not **redclare** variables, but the variables can be *shadowed*
*   All variables **MUST** be used, otherwise a *Compile Runtime* is created

*   Visibility
    *   lower case first letter for package scope
    *   upper case first letter to be used in a *GlobalScope*
    *   no private scope (all variables will be seen)

## Naming Conventions
*   PascalCase (ThisIsPascal) or camelCase (thisIsCamel)
*   Capitalize acronyms: HTTP, URL, etc
*   Keep the names of variables as short as reasonably possible; like a counter in a For-Loop
*   Use longer variable names for a *longer-life* if used frequently

## Type Conversions
*   destinationType(variable)
*   use **strconv** package for string conversion
*   Go does not do *implicit* type conversions; instead these have to be done **explicity**

## Primitive Types
### there are a number of different primitive types in Golang
*   Boolean type: **bool** the value is either: true or false; on or off; 0 or 1
    *   the default value of **bool** in Golang is *false* , or off or 0
