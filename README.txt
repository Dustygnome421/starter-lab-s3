CS4513: A Simple AWS S3 
==================================

Note, this document includes a number of design questions that can help your implementation. We highly recommend that you answer each design question **before** attempting the corresponding implementation.
These questions will help you design and plan your implementation and guide you towards the resources you need.
Finally, if you are unsure how to start the project, we recommend you visit office hours for some guidance on these questions before attempting to implement this project.


Team members
-----------------

1. Aman Gupta (agupta9@wpi.edu)
2. Oliver Reera (omreera@wpi.edu)

Design Questions
------------------

1. When implementing the `CreateBucket(bucketName string,)` function, you need to create the bucket in each node.
   1.1 How will you determine the number of nodes in the S3 system?
    Using the getNumberNodes() function. The common.go file keeps track of the number of nodes in the system.
   1.2 How will you create the directories?
    Using the mkdir() system call.

Brief response please!
---------------------

1. When implementing the `WriteNodeFile(nodeIndex int, bucketName string, fileName string, contents []byte, version time.Time,)` function, you need to write the file to the specified location.
   1.1 How will you find the relative path to the file?
    Create a pathname string, with the value <bucket>/<filename>.
   1.2 How will you handle if the file exists?
    The file will be overwritten, with all of its previous data deleted.
   1.3 How will you handle if the file doesn't exist?
    A new empty file will be created.
   1.4 How will you keep track of file versions?
    An extension will be concatenated to the filename. The version will be the time passed as an argument to the WriteNodeFile() function.

Brief response please!
---------------------

1. When implementing the `ReadNodeFile(nodeIndex int, bucketName string, fileName string)` function, you need to read the file at the specified location.
   1.1 How will you find the relative path to the file?
    The relative path will be nodes/<nodeIndex>/<bucketName>/<filename>.
   1.2 How will you get the version of that file?
    The version can be obtained from the ModTime() call on the file, which returns the last time the file was modified.

Brief response please!
---------------------

1. When implementing the `RequestWriteFile(bucketName string, fileName string, fileContents []byte)` function, you need to write the file to the specified bucket.
   1.1 How will you ensure that two clients accessing the same file at the same time will behave predictably? In other words, how will you handle scheduling requests?
    We can use a mutex to ensure that the first client who requests to write will be guaranteed to be the only one writing to it.
   1.2 How will you choose how many nodes to write to (hint: `src/replication/common.go`)?
    The quorum will need to be reached for nodes written to (majority of nodes).
   1.3 How will you choose which nodes to write to?
    We can write to the first n files, where n is the write quorum.
   1.4 How will you ensure that all copies of the file written in a single request have the same version?
    The version will be set to time.Now(), and then the version will be passed to the WriteNodeFile() function.

Brief response please!
---------------------

1. When implementing the `RequestReadFile(bucketName string, fileName string)` function, you need to read the file at the specified bucket.
   1.1 How will you ensure that two clients accessing the same file at the same time will behave predictably? In other words, how will you handle scheduling requests?
    Similarly to the RequestWriteFile(), the first client to obtain the mutex for the file will be able to read it, then release the mutex so that the next can read the file.
   1.2 How will you choose how many nodes to read from (hint: `src/replication/common.go`)?
    The request will read the majority of nodes, so that the quorum is reached.
   1.3 How will you choose which nodes to read from?
    The request can read from the first n nodes, where n is the read quorum.
   1.4 How will you determine which version of the file is the newest?
    All the contents of the files that are read are compared. If they are all the same, then the first file's contents are returned.
   1.5 How will you handle if a node cannot be read from (is faulty)?
    If the file's contents are nil, then the function skips over the individual file.

Brief response please!
---------------------

Errata
------
You can earn up to ten points for bugs, errors, and additional test cases. Two points each. 

6. Bug bounty: describe any known errors or bugs of the project writeup and starter code.
One bug we found had to do with when we added a check for the version of files when reading. In some cases, the versions of the files would be milliseconds off, leading to their versions being not equal, so haing a check to make sure their versions were the same lead to test cases not passing. 
Another bug we found was when all the test cases are run at once, every so often, the last test case would take longer and sometimes lead to a race condition. It does not happen every time though.

7. Additional Test Cases: describe any new test cases that you developed and explain why they are needed. 


---------------------

Misc 
-------
8. Describe any deviations, if any, from the requirement.
