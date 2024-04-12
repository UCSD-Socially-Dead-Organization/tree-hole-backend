# Treehole-backend Design Doc

## Problem Context
Nowadays, people are struggling with daily life, often seeking different types of emotional support. This product is one of them. It allows people to store their secrets and sadness in a treehole, similar to a diary, but no one knows who wrote it unless you choose to expose yourself. Then, every day, the system would assign a random user to visit the treehole and potentially provide some feedback. The intention of this product is to provide a social and positive environment for people who might be too shy to share their emotions and feelings, much like the author himself.

## Proposed Solution
1. The user can generate treehole messages, including a title and body. Additionally, users have the option to delete their treehole messages.

2. Users can create their own profiles, which include a username, age, friend list, and the option to upload a profile picture.

3. Users will receive a random treehole (potentially unique) from the system. They can enter the treehole to view all messages from the treehole owner. After viewing the messages, users can send a friend request along with an introductory message. If the treehole owner accepts the request, the user will be able to view their profile and exchange messages.

4. Users will have treehole mailboxes for messages from different users.

5. A report button to report disgusting users.

6. A user can also unfriend certain users.

## Goals and Non-Goals

填滿每個人的心跟樹洞

To Create a healthy social environment for this strange world.

### Non-Goals
This is not a dating app, so if a user gets reported many times by other users, their account will be banned.

## Design
TODO:
Three microservices will be utilized:
1.  Front end developed with React Native.
2.  Backend developed with Golang.
3.  Chronjob microservice (might be Golang as well) responsible for updating the current treehole for each user.
4.  The database will be PostgreSQL, chosen for its relational database capabilities to ensure adherence to the ACID properties.
5.  Deployment will be conducted using GCP (Google Cloud Platform).

## Alternatives Considered
PUB/SUB solution. People can subscribe to a certain type of treehole.

## Open Questions
PLEASE ASK IF YOU HAVE ANY QUESTIONS.</br>

## Parties Involved
Jeffrey Wu: the senior software engineer from TuSimple </br>
Jonathan Jiang: Grad school student from SJSU </br>
陳金華: Frontend tech lead </br>
### Reference docs
https://auth0.com/docs/quickstart/webapp/golang/interactive#install-dependencies
