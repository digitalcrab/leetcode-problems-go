# Isomorphic Strings

https://leetcode.com/problems/isomorphic-strings/description/

## Problem description

Given two strings `s` and `t`, determine if they are isomorphic.

Two strings `s` and `t` are isomorphic if the characters in `s` can be replaced to get `t`.

All occurrences of a character must be replaced with another character while preserving the order of characters. No two
characters may map to the same character, but a character may map to itself.

## Example 1:

```text
Input: s = "egg", t = "add"
Output: true
```

Explanation:
The strings s and t can be made identical by:
Mapping 'e' to 'a'.
Mapping 'g' to 'd'.

## Example 2:

```text
Input: s = "foo", t = "bar"
Output: false
```

Explanation:
The strings s and t can not be made identical as 'o' needs to be mapped to both 'a' and 'r'.

## Example 3:

```text
Input: s = "paper", t = "title"
Output: true
```

## Constraints:

```text
1 <= s.length <= 5 * 104
t.length == s.length
s and t consist of any valid ascii character.
```