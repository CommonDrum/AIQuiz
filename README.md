## What is it?
This is a commandline applicaiton that uses OpenAI API to generate quizes for specified topic. 

### How to run?
Before runnning the go project you need to add a system variable with you OpenAI API key.
Then just run the project as you would do with any other go program.

### Example
When prompted for specific subject program will generate a bunch of questions in this format:
```
{
      "topic": "Design Secure Architectures",
      "question": "A company stores sensitive user data in an Amazon S3 bucket. How can the Solutions Architect ensure that the data is encrypted at rest?",
      "options": {
        "A": "Enable Server-Side Encryption (SSE) on the S3 bucket",
        "B": "Enable Amazon S3 VPC Endpoint",
        "C": "Enable AWS WAF for the S3 bucket",
        "D": "Enable AWS Shield on the S3 bucket"
      },
      "answer": "A",
      "explanation": "Server-Side Encryption (SSE) in Amazon S3 encrypts data at rest and automatically decrypts it when it's retrieved, providing protection for the stored data."
    }
```

Then it will choose one of them at random and present you with it.

