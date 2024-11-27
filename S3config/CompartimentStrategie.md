    {
    "Version": "2012-10-17",
    "Statement": [
        {
            "Sid": "AllowSpecificOrigin",
            "Effect": "Allow",
            "Principal": "*",
            "Action": "s3:GetObject",
            "Resource": "arn:aws:s3:::studymove-cdn/*",
            "Condition": {
                "StringLike": {
                    "aws:Referer": [
                        "http://localhost:3000/*",
                        "https://votre-domaine.com/*"
                    ]
                }
            }
        }
    ]

}
