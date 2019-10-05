# dump

## Name

*dump* - dumps all incoming queries on standard output.

## Description

*dump* uses the synax from the *log* plugin, and defaults to this format:

~~~
{remote} - {>id} {type} {class} {name} {proto} {port}
~~~

So a query will show up as:

~~~
:1 - 3644 MX IN example.net. udp 46481
~~~

Note that this is shorter than the default for *log* so you can distinguish between the two outputs.
*log* only logs queries that have seen a response, so this plugin can be used as a debugging aid to
just dump all incoming queries.

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
