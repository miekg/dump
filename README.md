# dump

## Name

*dump* - dump all incoming queries on standard output.

## Description

*dump* uses the synax from the *log* plugin, and defaults to this format:

~~~
{remote} - [{when}] {>id} {type} {class} {name} {proto} {port}
~~~

So a query will show up as:

~~~
:1 - [17/Jan/2018:20:02:19 +0000] 3644 MX IN example.net. udp 46481
~~~

Note that this is shorter than the default for *log* so you can distinguish between the two outputs.

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
