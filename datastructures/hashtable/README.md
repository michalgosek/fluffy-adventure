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

When using open addressing as a collision resolution technique the **key-value pairs are stored in the table (array)
itself** as opposed to a data structure like in separate chaining. 

This means we need to care a great deal about the size of our hash table and how many elements are currently
in the table.

Load factor = items in table / size of table 

The O(1) constant time behaviour attributed to hash tables assumes the load factor (α) is kept below a certain
fixed value. This meas once α > **threshold** we need to grow the table size (ideally exponentially, e.g double). 

When we want to insert a key-value pair (k, v) into the hash table we hash the key and obtain an original position 
for where this key-value pair belongs, i.e **H**(k).

If the position our key hashed to is occupied we try another position in the hash table by offsetting the current 
position subject to a **probing sequence P**(x). We keep doing this until an unoccupied slot is found. 

There are an infinite amount of probing sequences you can come up with, here a few:

- Linear probing:
**P**(x) = ax + b, where a, b are constants
- Quadratic probing:
**P**(x) = ax^2 + bx + c, where a,b,c are constants 
- Double hashing:
**P**(k, x) = x * **H2**(k), where **H2**(k) is a secondary hash function
- Pseudo random number generator:\
**P**(k, x) = x * **RNG**(**H**(K), x), where RNG is random number generator function seeded with **H**(k)

General insertion method for open addressing on a table of size N goes as follows:

```
x := 1 
keyHash := H(k)

while table[index] != null:
    index = (keyHash + P(k, x)) mod N 
    x = x + 1 

insert (k, v) at table index
```

Where **H**(k) is the hash for the key k and **P**(k, x) is the probing function 

Prob function is always setup such way it returns the unoccupied index position because the load factor is 
below threshold. 

Most randomly selected probing sequences modulo **N** will produce a cycle shorter than the table size. 

This becomes problematic when you are trying to insert a key-value pair and all the buckets on the cycle are 
occupied because you will get struck in an **infinite loop**! 

Q: So that's concerting how do we handle probing functions which produce cycles shorter than the table size?

A: _In general the consensus is that we don't handle this issue, instead we avoid it altogether by restricting 
our domain of probing functions to those which produce a cycle of exactly length N*._ 

There are few exceptions with special properties that can produce shorter cycles.

Techniques such as linear probing, quadratic probing and double hashing are all subject to the issue of causing 
cycles which is why the probing functions used with these methods are very specific. 

Notice that open addressing is very sensitive to the hashing function and probing function used. This is not something
you have to worry about (as much) if you are using **separate chaining** as collision resolution method. 

Q: So how do we pick a probing function we can work with?

A: _There are numerous ways, but three of the most popular approaches are_:

1) Let P(x) = x^2, keep the table size a prime number > 3 and also keep (alfa) <= 0.5
2) Let P(x) = (x^2 + x) / 2 and keep the table size a power of two
3) Let P(x) = (-1)*x^2 and keep the table size a prime N where N congruent 3 mod 4 

# Linear probing 

Linear probing is **probing method** which probes according to a linear formula, specifically:\
P(x) = ax + b  where a != 0, b are constants \
(Note: The constant b is obsolete)

Q: Which value(s) of the constant a in P(x) = ax produce a full cycle modulo N?
A: This happens when a and N are relatively prime. Two numbers are relatively prime if their Greatest Common
Denominator (GCD) is equal to one. Hence, when GCD(a, n) = 1 the probing function P(x) be able to generate
a complete cycle, and we will always be able to find an empty bucket! 

# Quadratic probing 
Quadratic probing is a **probing method** which probes according to a quadratic formula, specifically:
P(x) = ax^2 + bx + c, where a,b,c are constants and a != 0 (otherwise we have linear probing)

# Double hashing 
Double hashing is a **probing method** which probes according to a constant multiple of another hash function, 
specifically:

**P**(k, x) = x * **H2**(k), where **H2**(k) is a second hash function. 

**H2**(k) must hash the same type of keys as **H1**(k)

To fix the issue of cycles pick the table size to be a prime number and also compute the value of 
delta. 

δ = H2(k) mod N

If δ = 0 then we are guaranteed to be stuck in a cycle, so when this happens set δ = 1 

Notice that 1 <= δ < N and **GCD**(δ, N) = 1 since N is prime. Hence, with these conditions 
we know that modulo N the sequence. 

H1(k), H1(k)+1δ, H1(k)+2δ, H1(k)+3δ, H1(k)+4δ, ... 
is certain to have order N 

### Constructing H2(k)

Suppose the key k has type T

Whenever we want to use double hashing as a collision resolution method we need to fabricate
a new function **H2**(k) that knows how to hash keys of type T.

It would be nice to have a systematic way to be able to effectively produce a new hash function 
every time we need one, right?

Luckily for us the keys we need to hash are always composed of the same fundamental building blocks.
In particular: integers, strings, real numbers, fixed length vectors, etc...

There are many well known high quality hash functions for these fundamental data types. Hence, we can use and combine
them to construct our function **H2**(k).

Frequently the hash functions selected to compose **H2**(k) are picked from a pool of hash functions called **universal 
hash functions** which generally operate on once fundamental data type. 
