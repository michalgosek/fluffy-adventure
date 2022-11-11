# What is a Hash table?
A **hash table (HT)** is a data structure that provides a mapping from keys to value 
using a technique called hashing.

| Key     | Value |
|---------|-------|
| Michael | Green |
 
We refer to these as key-value pairs. Keys must be unique, but values can be repeated.

HTs are often used to track item frequencies. For instance, counting the number of times
a word appears in a given text.

The key-value pairs you can place in a HT can be of any type not just strings and numbers,
but also objects! However, the keys needs to be **hashable**, a property we will discuss shortly. 

To be able to understand how a mapping is constructed between key-value pairs we first need 
to talk about hash functions. 

A **hash function H(x)** is a function that maps a key 'x' to a whole number in a fixed range. 

For example, H(x) = (x^2 - 6x + 9) mod 10 maps all integer keys to the range [0, 9]

* H(4) = (16 - 24 + 9) mod 10 = 1
* H(-7) = (49 + 42 + 9) mod 10 = 0
* H(0) = (0 - 0 + 9) mod 10 = 9
* H(2) = (4 - 12 + 9) mod 10 = 1
* H(8) = (64 - 48 + 9) mod 10 = 5  

Output is not unique, there can be two values that passed into hash function maps to the same output.
ex. H(4) and H(2) returns 1. What's totally fine.  

We can also define hash functions for arbitrary objects such as strings, lists, tuples, multi data objects, etc.

If **H(x) = H(y)** then objects x and y **might be equal**, but if **H(x) != H(y)** the x and y are **certainly not equal**. 

Q: How we can use this to our advantage to speedup object comparisons?\
A: _This mean that instead of comparing x and y directly a smarter approach is to first compare their hash values, 
and only if the hash values match do we need to explicitly compare x and y._ 

A hash function H(x) must be deterministic. This means that if H(x) = y then H(x) must always produce y and 
never another value. This may seem obvious, but it is critical to the functionality of a hash function. 

We try very hard to make **uniform** hash functions to minimize the number of the hash collisions.

A **hash collision** is when two objects x, y hash to the same value (i.e. H(x) = H(y))
 
Q: What makes a key of type T hashtable?\
A: _Since we are going to use hash functions in the implementation of the hash table we need our 
hash functions to be deterministic. To enforce this behaviour, we demand that the **keys used in our
hash table are immutable data types (ex. string, number literals opposite to sets, lists etc.).**
Hence, if a key of type T is immutable, and we have a hash function H(k) defined for all keys k of type T_

#  How does a hash table work?

Ideally we would like to have a very fast insertion, lookup and removal time for the data we are placing 
withing our hash table.  Remarkably, we can achieve all this in O(1)* time using a **hash function as way to index into a hash table**. 

*The constant time behaviour attributed to hash tables is only true if you have a good **uniform hash function**. 
 
Think of the hash table below as indexable block of memory (an array) and we can only access its entries using 
the value given to us by our hash function **H**(x).


|     | Key | Value |
|-----|-----|-------|
| 1   |     |       |
| 2   |     |       |
| 3   |     |       |
| 4   |     |       |
| 5   |     |       |
| 6   |     |       |


Suppose we're inserting (integer, string) key-value pairs into the table representing rankings of users to their 
usernames from an online programming competition and we're using the hash function:

**H**(x) = x^2 + 3 mod 10 

To **insert** (3, "byte-eater") we hash the key (the rank) and find out where it goes in the table. 

* H(3) = (3^2 + 3) mod 10 = 2 
* H(1) = (1^2 + 3) mod 10 = 4 
* H(0) = (0^2 + 3) mod 10 = 3 

|     | Key | Value        |
|-----|-----|--------------|
| 0   |     |              |
| 1   |     |              |
| 2   | 3   | "byte-eater" |
| 3   | 0   | "jerry33"    |
| 4   | 1   | "migos"      |
| 5   |     |              |
| 6   |     |              |

To lookup which user has rank **r** we simply compute **H**(r) and look inside the hashtable. 

Q: What do we do if there is a hash collision?
For example, users with rank 2 and 8 hash to the same value.

- H(2) = 2^2 + 3 mod 10 = **7**
- H(8) = 2 ^8 + 3 mod 10 = **7**

A: _We use one of many hash collision resolution techniques to handle this, the two most popular ones are **separate 
chaining** and **open addressing**._ 

**Complexity Analysis**:

| Operation | Avg   | Worst |
|-----------|-------|-------|
| Search    | O(1)* | O(n)  |
| Insertion | O(1)* | O(n)  |
| Removal   | O(1)* | O(n)  |

*The constant time behaviour attributed to hash tables is only true if you have a good uniform hash function!

# What is Separate Chaining? 

**Separate chaining** deals with hash collisions by maintaining a data structure (usually a linked list) to hold all 
the different values (key-value pairs) which hashed to a particular value. 

Note: The data structure used to cache the items which hashed to a particular value is not limited to a linked list.
Some implementations use one or a mixture of: arrays, binary trees, self-balancing trees etc. 

It may happen that the value you are looking for does not exist in the bucket the key hashed to in which case 
the item does not exist in the hash table. 

Q: How do I maintain **O(1)** insertion and lookup time complexity once my hash table gets really full, 
and I have long linked  list chains?

A: _Once the hash table contains a lot of elements you should create a new hash table with a larger capacity 
and rehash all the items inside old hash table and disperse them throughout the new hash table at different
locations._ 


Q: How do I **remove** key-value paris from my hash table?

A: _Apply the same procedure as doing lookup for a key, but this time instead of returning the value associated 
with the key remove the node in the linked list data structure._ 

Q: Can I use another data structure to model the bucket behaviour required for the separate chaining method?

A: _Of course! Common data structures used instead of a linked list include: arrays, binary trees, self-balancing 
trees etc. You can even go with a hybrid approach like Java's HashMap. However, note that some of these are much more
memory intensive and complex to implement than a simple linked list which is why they may be less popular._ 


# What is Open Addressing?

**Open addressing** deals with hash collisions by finding another place within the hash table for the object to go by 
offsetting it from the position to which it hashed to.
