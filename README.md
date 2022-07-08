# POC Serverless + AWS + GO + Segment

## :bulb: Prerequisites

* [Node.js](https://nodejs.org/en/) can be installed using [nvm](https://github.com/nvm-sh/nvm)
* [Serverless](https://www.serverless.com/) installed locally : `npm i -g serverless`
* an AWS account : [AWS Console](https://aws.amazon.com/)
* a Segment account : [Segment](https://segment.com/)

## :floppy_disk: Configuration

* install the project `npm install`
* Update `.env` file and add fill the `SEGMENT_TOKEN` variable with your token. 
* [Create an aws access key](https://aws.amazon.com/premiumsupport/knowledge-center/create-access-key/)
* Configure your aws credentials locally : 
    * using aws cli : https://docs.aws.amazon.com/cli/latest/userguide/cli-configure-files.html
    * using serverless : `serverless config credentials --provider aws --key <key> --secret <secret>`

## :rocket: Deploy the application

`make deploy`

It will build the application and push it on AWS using serverless (see [Makefile](./Makefile))

## :clapper: Use Case

The use case is about to consume SQS messages in order to send Segment's notifications.

Once the application is deployed on AWS, use the console in order to manually post a message into the SQS queue with this content :
```
{
    "EventName": "Awesome event",
    "userId": 1234
}
```