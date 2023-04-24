CS4513: A Simple AWS S3 
==================================

Note, this document includes a number of design questions that can help your implementation. We highly recommend that you answer each design question **before** attempting the corresponding implementation.
These questions will help you design and plan your implementation and guide you towards the resources you need.
Finally, if you are unsure how to start the project, we recommend you visit office hours for some guidance on these questions before attempting to implement this project.


Team members
-----------------

1. Alice (alice@wpi.edu)
2. Bob (bob@wpi.edu)

Design Questions
------------------

1. When implementing the `CreateBucket(bucketName string,)` function, you need to create the bucket in each node.
   1.1 How will you determine the number of nodes in the S3 system?
   1.2 How will you create the directories?

Brief response please!
---------------------

1. When implementing the `WriteNodeFile(nodeIndex int, bucketName string, fileName string, contents []byte, version time.Time,)` function, you need to write the file to the specified location.
   1.1 How will you find the relative path to the file?
   1.2 How will you handle if the file exists?
   1.3 How will you handle if the file doesn't exist?
   1.4 How will you keep track of file versions?

Brief response please!
---------------------

1. When implementing the `ReadNodeFile(nodeIndex int, bucketName string, fileName string)` function, you need to read the file at the specified location.
   1.1 How will you find the relative path to the file?
   1.2 How will you get the version of that file?

Brief response please!
---------------------

1. When implementing the `RequestWriteFile(bucketName string, fileName string, fileContents []byte)` function, you need to write the file to the specified bucket.
   1.1 How will you ensure that two clients accessing the same file at the same time will behave predictably? In other words, how will you handle scheduling requests?
   1.2 How will you choose how many nodes to write to (hint: `src/replication/common.go`)?
   1.3 How will you choose which nodes to write to?
   1.4 How will you ensure that all copies of the file written in a single request have the same version?

Brief response please!
---------------------

1. When implementing the `RequestReadFile(bucketName string, fileName string)` function, you need to read the file at the specified bucket.
   1.1 How will you ensure that two clients accessing the same file at the same time will behave predictably? In other words, how will you handle scheduling requests?
   1.2 How will you choose how many nodes to read from (hint: `src/replication/common.go`)?
   1.3 How will you choose which nodes to read from?
   1.4 How will you determine which version of the file is the newest?
   1.5 How will you handle if a node cannot be read from (is faulty)?

Brief response please!
---------------------

Errata
------
You can earn up to ten points for bugs, errors, and additional test cases. Two points each. 

6. Bug bounty: describe any known errors or bugs of the project writeup and starter code. 


7. Additional Test Cases: describe any new test cases that you developed and explain why they are needed. 


---------------------

Misc 
-------
8. Describe any deviations, if any, from the requirement.
