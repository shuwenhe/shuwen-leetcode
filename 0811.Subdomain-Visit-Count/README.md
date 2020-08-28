# [811. Subdomain Visit Count](https://leetcode.com/problems/subdomain-visit-count/)


## 题目

A website domain like "discuss.leetcode.com" consists of various subdomains. At the top level, we have "com", at the next level, we have "leetcode.com", and at the lowest level, "discuss.leetcode.com". When we visit a domain like "discuss.leetcode.com", we will also visit the parent domains "leetcode.com" and "com" implicitly.

Now, call a "count-paired domain" to be a count (representing the number of visits this domain received), followed by a space, followed by the address. An example of a count-paired domain might be "9001 discuss.leetcode.com".

We are given a list `cpdomains` of count-paired domains. We would like a list of count-paired domains, (in the same format as the input, and in any order), that explicitly counts the number of visits to each subdomain.

    Example 1:
    Input: 
    ["9001 discuss.leetcode.com"]
    Output: 
    ["9001 discuss.leetcode.com", "9001 leetcode.com", "9001 com"]
    Explanation: 
    We only have one website domain: "discuss.leetcode.com". As discussed above, the subdomain "leetcode.com" and "com" will also be visited. So they will all be visited 9001 times.

    Example 2:
    Input: 
    ["900 google.mail.com", "50 yahoo.com", "1 intel.mail.com", "5 wiki.org"]
    Output: 
    ["901 mail.com","50 yahoo.com","900 google.mail.com","5 wiki.org","5 org","1 intel.mail.com","951 com"]
    Explanation: 
    We will visit "google.mail.com" 900 times, "yahoo.com" 50 times, "intel.mail.com" once and "wiki.org" 5 times. For the subdomains, we will visit "mail.com" 900 + 1 = 901 times, "com" 900 + 50 + 1 = 951 times, and "org" 5 times.

**Notes:**

- The length of `cpdomains` will not exceed `100`.
- The length of each domain name will not exceed `100`.
- Each address will have either 1 or 2 "." characters.
- The input count in any count-paired domain will not exceed `10000`.
- The answer output can be returned in any order.


## 题目大意


一个网站域名，如 "discuss.leetcode.com"，包含了多个子域名。作为顶级域名，常用的有 "com"，下一级则有 "leetcode.com"，最低的一级为 "discuss.leetcode.com"。当我们访问域名 "discuss.leetcode.com" 时，也同时访问了其父域名 "leetcode.com" 以及顶级域名 "com"。给定一个带访问次数和域名的组合，要求分别计算每个域名被访问的次数。其格式为访问次数+空格+地址，例如："9001 discuss.leetcode.com"。

接下来会给出一组访问次数和域名组合的列表 cpdomains 。要求解析出所有域名的访问次数，输出格式和输入格式相同，不限定先后顺序。



## 解题思路


- 这一题是简单题，统计每个 domain 的出现频次。每个域名根据层级，一级一级的累加频次，比如 `discuss.leetcode.com`、`discuss.leetcode.com` 这个域名频次为 1，`leetcode.com` 这个域名频次为 1，`com` 这个域名频次为 1。用 map 依次统计每个 domain 出现的频次，按照格式要求输出。
