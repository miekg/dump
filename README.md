# dump

## Name

*dump* - dump all incoming queries on standard output.

## Description

*dump* uses the synax from the *log* plugin, and defaults to this format:

~~~
{remote} - [{when}] {>id} {type} {class} {name} {proto} {port}
~~~

So each query will show up as:

## Syntax

~~~ txt
dump
~~~

## Examples

Dump all queries.

~~~ corefile
. {
    dump
}
~~~
