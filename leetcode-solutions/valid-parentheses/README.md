# Valid Parentheses

https://leetcode.com/problems/valid-parentheses/description/

## Problem Description

Given a string `s` containing just the characters `'('`, `')'`, `'{'`, `'}'`, `'['` and `']'`, determine if the input
string is valid.

An input string is valid if:

1. Open brackets must be closed by the same type of brackets.
2. Open brackets must be closed in the correct order.
3. Every close bracket has a corresponding open bracket of the same type.

## Example 1

```text
Input: s = "()"
Output: true
```

## Example 2

```text
Input: s = "()[]{}"
Output: true
```

## Example 3

```text
Input: s = "(]"
Output: false
```

## Example 4

```text
Input: s = "([])"
Output: true
```

## Constraints

1. 1 <= s.length <= 104
2. `s` consists of parentheses only `'()[]{}'`.
