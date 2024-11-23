# Gobyd - A simple bidding REST API

This is a simple bidding REST API that allows users to create auctions and place bids on them.
The API is build using mutual exclusion to ensure that only one bid is accepted at a time.
We also use replication to ensure that the data is not lost in case of a failure on a node.

## Vector Clocks

What is vector clocks?

## Mutual exclusion

Algorithm: Ricart & Agrawala

## Replication

Active vs Passive replication
