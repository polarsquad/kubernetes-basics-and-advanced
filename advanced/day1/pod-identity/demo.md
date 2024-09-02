# Pod Identity DEMO

Requirements:  
- working cluster, at least version 1.24

## Create IAM Role for Pod Identity

Role:  
```
{
	"Version": "2012-10-17",
	"Statement": [
		{
			"Action": [
				"s3:ListBucket",
				"s3:GetBucketLocation"
			],
			"Effect": "Allow",
			"Resource": [
				"arn:aws:s3:::stetson-bucket"
			]
		}
	]
}
```

Trust Policy:
```
{
    "Version": "2012-10-17",
    "Statement": [
        {
            "Effect": "Allow",
            "Principal": {
                "Service": "pods.eks.amazonaws.com"
            },
            "Action": [
                "sts:AssumeRole",
                "sts:TagSession"
            ],
            "Condition": {
                "ArnEquals": {
                    "aws:SourceArn": [
                        "arn:aws:eks:eu-north-1:050548751257:cluster/stetson"
                    ]
                }
            }
        }
    ]
}
```

