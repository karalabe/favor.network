// Copyright (C) 2017 Péter Szilágyi - All Rights Reserved.
//
// This file is part of the `favornet` Status IM hackathon entry. Source code
// is provided solely for judging purposes, without any explicit or implied
// waiver of rights or allowance of reuse in whole or in part.
pragma solidity ^0.4.11;

// Favornet is the favor-network's ledger contract that keeps track of favors
// users requested from one another, promises made to repay them and current
// balances of who is "owed one" by whom and why.
contract Favornet {
    // Request is a currently unfulfilled favor request. Once a request is
    // accepted, the only way to get rid of it is to fulfil it, issuing a
    // promise to repay the favor (or rewarding with an existing promise).
    struct Request {
        uint     Id;     // Unique identifier of this favor request
        address  From;   // Address of the user to fulfill the favor
        string   Favor;  // Description of the favor being asked
        bool     Bound;  // Whether the request was accepted to fulfil
        uint     Reward; // Promise to reward in exchange for this favor
    }
    // Promise is a pending commitment issued by a user for someone doing a
    // favor for her. Promises may be swapped or traded at will, eventually
    // reaching back to the indebted party for honouring it via a favor.
    struct Promise {
        string  Favor;   // Description of the favor promised to be repayed
        address From;    // Address of the user making the promise to fulfil
        address Owner;   // Current owner of the favor return commitment
        uint    Offered; // Favor request offering this promise as reward
    }

    uint autoid;                        // Next id to assign to a new request
    mapping (uint => Request) requests; // Set of unfulfilled favor requests
    mapping (uint => Promise) promises; // Set of pending return commitments

    mapping (address => uint[]) liabilities; // Favors requested by users
    mapping (address => uint[]) assets;      // Commitments owned by users

    // Favornet is the contract constructor. For now it only initializes the
    // autoid to 1 to reserve 0 as an unusable special case.
    function Favornet() {
        autoid = 1;
    }
    // MakeRequest creates a new favor request from the specified user. The
    // favor by default rewards with a new promise to return the favor, but
    // the user may offer for exchange an existing promise too instead.
    function MakeRequest(address from, string favor, uint reward) {
        // If the user offered an existing promise as a reward, ensure it's
        // actually owned by the sender and that it's not locked up in some
        // other already pending request.
        if (reward != 0) {
            Promise promise = promises[reward];
            if (promise.Owner != msg.sender || promise.Offered != 0) {
                throw;
            }
            promise.Offered = autoid;
        }
        // Reward checks out, create the request and return
        liabilities[msg.sender].push(autoid);
        requests[autoid] = Request({
            Id:     autoid,
            Favor:  favor,
            From:   from,
            Bound:  false,
            Reward: reward
        });
        autoid++;
    }
    // DropRequest deletes an unfulfilled and unboud favor request from the users
    // own list of liabilities. Requests that have already been accepted cannot
    // be dropped.
    //
    // The reason both index and id are required for this operation is to allow
    // calculating the index off chain (cheap), but still retain the guarantee
    // of operating on the correct request object.
    function DropRequest(uint index, uint id) {
        // Ensure the request exists and is still unbound
        uint count = liabilities[msg.sender].length;
        if (index >= count) {
            throw;
        }
        if (liabilities[msg.sender][index] != id) {
            throw;
        }
        if (requests[id].Bound) {
            throw;
        }
        // All checks passed, get rid of the favor request
        liabilities[msg.sender][index] = liabilities[msg.sender][count-1];
        liabilities[msg.sender].length--;
    }
    // AcceptRequest makes a binding agreement to honour a given favor request.
    // From this point onward the request can only ever be deleted if its author
    // honors it with a promise reward, otherwise
    //
    // The reason both index and id are required for this operation is to allow
    // calculating the index off chain (cheap), but still retain the guarantee
    // of operating on the correct request object.
    function AcceptRequest(address from, uint index, uint id) {
      // Ensure the request exists, is for us and still unbound
      uint count = liabilities[from].length;
      if (index >= count) {
          throw;
      }
      if (liabilities[from][index] != id) {
          throw;
      }
      if (requests[id].From != msg.sender) {
          throw;
      }
      if (requests[id].Bound) {
          throw;
      }
      // Everything checks out, make a binding commitment
      requests[id].Bound = true;
    }

    // GetRequestCount return the number of favor requests a user currently has open.
    function GetRequestCount(address user) constant returns (uint) {
        return liabilities[user].length;
    }
    // GetRequestAt returns the currently unfulfilled favor requests of a particular
    // user at a particular storage slot.
    function GetRequestAt(address user, uint index) constant returns (uint id, address from, string favor, bool bound, uint reward) {
        Request req = requests[liabilities[user][index]];
        return (req.Id, req.From, req.Favor, req.Bound, req.Reward);
    }
    // GetRequest returns the specified currently unfulfilled favor requests.
    function GetRequest(uint id) constant returns (address from, string favor, bool bound, uint reward) {
        Request req = requests[id];
        return (req.From, req.Favor, req.Bound, req.Reward);
    }
}
