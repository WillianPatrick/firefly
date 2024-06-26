---
title: API Query Syntax
---

## Syntax Overview

REST collections provide filter, `skip`, `limit` and `sort` support.

- The field in the message is used as the query parameter
  - Syntax: `field=[modifiers][operator]match-string`
- When multiple query parameters are supplied these are combined with AND
- When the same query parameter is supplied multiple times, these are combined with OR

## Example API Call

`GET` `/api/v1/messages?confirmed=>0&type=broadcast&topic=t1&topic=t2&context=@someprefix&sort=sequence&descending&skip=100&limit=50`

This states:

- Filter on `confirmed` greater than 0
- Filter on `type` exactly equal to `broadcast`
- Filter on `topic` exactly equal to `t1` _or_ `t2`
- Filter on `context` containing the case-sensitive string `someprefix`
- Sort on `sequence` in `descending` order
- Paginate with `limit` of `50` and `skip` of `100` (e.g. get page 3, with 50/page)

Table of filter operations, which must be the first character of the query string (after the `=` in the above URL path example)

### Operators

Operators are a type of comparison operation to
perform against the match string.

| Operator | Description             |
| -------- | ----------------------- |
| `=`      | Equal                   |
| (none)   | Equal (shortcut)        |
| `@`      | Containing              |
| `^`      | Starts with             |
| `$`      | Ends with               |
| `<<`     | Less than               |
| `<`      | Less than (shortcut)    |
| `<=`     | Less than or equal      |
| `>>`     | Greater than            |
| `>`      | Greater than (shortcut) |
| `>=`     | Greater than or equal   |

> Shortcuts are only safe to use when your match
> string starts with `a-z`, `A-Z`, `0-9`, `-` or `_`.

### Modifiers

Modifiers can appear before the operator, to change its
behavior.

| Modifier | Description                                    |
| -------- | ---------------------------------------------- |
| `!`      | Not - negates the match                        |
| `:`      | Case insensitive                               |
| `?`      | Treat empty match string as null               |
| `[`      | Combine using `AND` on the same field          |
| `]`      | Combine using `OR` on the same field (default) |

## Detailed examples

| Example    | Description                               |
| ---------- | ----------------------------------------- |
| `cat`      | Equals "cat"                              |
| `=cat`     | Equals "cat" (same)                       |
| `!=cat`    | Not equal to "cat"                        |
| `:=cat`    | Equal to "CAT", "cat", "CaT etc.          |
| `!:cat`    | Not equal to "CAT", "cat", "CaT etc.      |
| `=!cat`    | Equal to "!cat" (! is after operator)     |
| `^cats/`   | Starts with "cats/"                       |
| `$_cat`    | Ends with with "\_cat"                    |
| `!:^cats/` | Does not start with "cats/", "CATs/" etc. |
| `!$-cat`   | Does not end with "-cat"                  |
| `?=`       | Is null                                   |
| `!?=`      | Is not null                               |

## Time range example

For this case we need to combine multiple queries on the same `created`
field using AND semantics (with the `[`) modifier:

```
?created=[>>2021-01-01T00:00:00Z&created=[<=2021-01-02T00:00:00Z
```

So this means:

- `created` greater than `2021-01-01T00:00:00Z`
- `AND`
- `created` less than or equal to `2021-01-02T00:00:00Z`
