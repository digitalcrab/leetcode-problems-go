# Ransom Note

https://leetcode.com/problems/ransom-note/description/

## Problem description

Given two strings `ransomNote` and `magazine`, return `true` if `ransomNote` can be constructed by using the letters
from `magazine` and `false` otherwise.

Each letter in `magazine` can only be used once in `ransomNote`.

## Example 1:

```text
Input: ransomNote = "a", magazine = "b"
Output: false
```

## Example 2:

```text
Input: ransomNote = "aa", magazine = "ab"
Output: false
```

## Example 3:

```text
Input: ransomNote = "aa", magazine = "aab"
Output: true
```

## Constraints:

```text
1 <= ransomNote.length, magazine.length <= 105
ransomNote and magazine consist of lowercase English letters.
```