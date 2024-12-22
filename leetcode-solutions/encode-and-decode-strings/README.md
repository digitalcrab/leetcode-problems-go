# Encode and Decode Strings

https://leetcode.com/problems/encode-and-decode-strings/description/

## Problem Description

Design an algorithm to encode a list of strings to a string.
The encoded string is then sent over the network and is decoded back to the original list of strings.

## Example 1

```text
Input: dummy_input = ["Hello","World"]
Output: ["Hello","World"]
Explanation:
Machine 1:
Codec encoder = new Codec();
String msg = encoder.encode(strs);
Machine 1 ---msg---> Machine 2

Machine 2:
Codec decoder = new Codec();
String[] strs = decoder.decode(msg);
```

## Example 1

```text
Input: dummy_input = [""]
Output: [""]
```

## Constraints

```text
1 <= strs.length <= 200
0 <= strs[i].length <= 200
strs[i] contains any possible characters out of 256 valid ASCII characters.
```

## Follow up

Could you write a generalized algorithm to work on any possible set of characters?
