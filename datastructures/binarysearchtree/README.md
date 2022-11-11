# What is a Tree?
A tree is an undirected graph which satisfies any of the following definition:
 
1) An acyclic connected graph
2) A connected graph with n nodes and n-1 edges 
3) A graph in which any two vertices are connected by exactly one path. 

**Root node** is a top most node.
**A leaf node** is a node with no children.  

If we have a rooted tree then we will want to have a reference to the root node 
of our tree. It does not always matter which node is selected to be the root node because any 
node can become the root. 

A child is a node extending from another node. A parent is the inverse of this. 

Q: *What is the parent of the root node?*\
It has no parent, although it may be useful to assign the parent of the root node to be itself.
(e.g. filesystem tree)

A subtree is a tree entirely contained within another. They are usually denoted using triangles.\
Note: Subtrees may consist of a single node! 

# What is a Binary Tree?
A binary tree is a tree for which every node has at most two child nodes.

# What is a Binary Search Tree?

A binary search tree is a binary tree that satisfies the BST invariant:
- left subtree has smaller elements 
- right subtree has larger elements 

# When and where are Binary Search Tree? used?

1) Implementation of some map and set Abstract Data Types
   1) Red black Trees
   2) AVL Trees
   3) Splay Trees
   4) etc...
2) Used in the implementation of binary heaps 
3) Syntax trees (used by compiler and calculators)
4) Treap - a probabilistic DS (used as randomized BST)


# Adding element to a BST 

Binary Search Tree (BST) elements must be comparable so that we can 
order them inside the tree. 

When inserting an element we want to compare its value to the value 
stored in the current node we're considering to decide on one of the following:

- Recurse down to left subtree ( < case )
- Recurse down to right subtree ( > case )
- Handle fining a duplicate value ( = case)
- Create a new node (found a null leaf)

# Removing elements from a BST

Removing elements from a Binary Search Tree (BST) can be seen as two step process.

1) **Find** the element we wish to remove (if it exists)
2) **Replace** the node we want to remove with its successor (if any) to maintain
   the BST invariant.

Recall the _BST invariant_: left subtree has smaller elements and right subtree has larger
elements. 


**Find phase**

When searching our BST for a node with particular value one of four things will happen:
1) We hit a **null node** at which point we know value does not exist within our BST
2) Comparator value **equal to 0** (found it!)
3) Comparator value **less than 0** (the value if it exists, is in the left subtree)
4) Comparator value **greater than 0** (the value, if it exists, is in the right subtree)


**Remove phase**

1) Node to remove is a leaf node (no subtrees)
2) Node to remove has a right subtree but not left subtree 
3) Node to remove has a left subtree but not right subtree
4) Node to remove has both a left subtree and a right subtree

_case IV:_ 

_Q_: In which subtree will the successor of the node we are trying to remove be?\
_A:_ The answer is both! The successor can either be the **largest value** in the left 
subtree OR the **smallest value** in the **right subtree**. 

A justification for why there could be more than one successor is:

The **largest value** in the **left subtree** satisfies the BST invariant since it:

1) Is larger than everything in left subtree. This follows immediately from the definition 
of being largest. 
2) Is smaller than everything in right subtree because it was found in the left subtree.
 

The **smallest value** in the **right subtree** satisfies the BST invariant since:
1) Is smaller than everything in right subtree. This follows immediately from the definition of being the smallest.
2) Is larger than everything in left subtree because it was found in the right subtree.

# Tree Traversals (Preoder, Inorder, Postorder & Level Order)

1. Preorder - prints the value of the node **before** recursive calls
2. Inorder - prints the value of the node **between** the recursive calls 
3. Postorder - prints the value of the node **after** recursive calls 

Levelorder - prints the value of the nodes being on the same layer.
To obtain this ordering we want to do a **Breadth First Search** (BFS) from the root node down to the leaf nodes.

To do BFS we will need to maintain a **Queue** of the nodes left to explore. 
Begin with the root inside the queue and finish when the queue is empty. 

**Complexity Analysis**:

| Operation | Avg Time  | Worst |
|-----------|-----------|-------|
| Insert    | O(log(n)  | O(n)  |
| Delete    | O(log(n)  | O(n)  |
| Remove    | O(log(n)  | O(n)  | 
| Search    | O(log(n)) | O(n)  |

In worst case all operations have O(n) time complexity. 
