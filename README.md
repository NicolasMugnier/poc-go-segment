# POC Serverless + AWS + GO + Segment

## :bulb: Prerequisite

* an AWS account : [AWS Console](https://aws.amazon.com/)
* a Segment account : [Segment](https://segment.com/)

## :floppy_disk: Configuration

* Update `.env` file and add fill the `SEGMENT_TOKEN` variable with your token. 
* [Create an aws access key](https://aws.amazon.com/premiumsupport/knowledge-center/create-access-key/)
* Configure your aws credentials locally : 
    * using aws cli : https://docs.aws.amazon.com/cli/latest/userguide/cli-configure-files.html
    * using serverless : `serverless config credentials --provider aws --key <key> --secret <secret>`

## :rocket: Deploy the application

`npm install`

`make deploy`

## :clapper: Use Case

The use case is about consume SQS messages in order to send Segment's notifications.