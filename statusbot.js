// Copyright (C) 2017 Péter Szilágyi - All Rights Reserved.
//
// This file is part of the `favor.network` Status IM hackathon entry. Source code
// is provided solely for judging purposes, without any explicit or implied waiver
// of rights or allowance of reuse in whole or in part.

// favornetAddress is the Favor Network's contract address deployed on the Ropsten
// test network.
var favornetAddress = "0xaE287510a5b2E1ecd9538e75DC0Beff1B0e671c8";

// favornetABI is the Favor Network's contract ABI to allow user interaction and
// data queries.
var favornetABI = [{"constant":false,"inputs":[{"name":"index","type":"uint256"},{"name":"id","type":"uint256"}],"name":"DropPromise","outputs":[],"payable":false,"type":"function"},{"constant":true,"inputs":[{"name":"user","type":"address"},{"name":"index","type":"uint256"}],"name":"GetRequestAt","outputs":[{"name":"id","type":"uint256"},{"name":"from","type":"address"},{"name":"favor","type":"string"},{"name":"bound","type":"bool"},{"name":"reward","type":"uint256"}],"payable":false,"type":"function"},{"constant":false,"inputs":[{"name":"from","type":"address"},{"name":"index","type":"uint256"},{"name":"id","type":"uint256"}],"name":"AcceptRequest","outputs":[],"payable":false,"type":"function"},{"constant":false,"inputs":[{"name":"reqIdx","type":"uint256"},{"name":"reqId","type":"uint256"},{"name":"promIdx","type":"uint256"}],"name":"HonourRequest","outputs":[],"payable":false,"type":"function"},{"constant":false,"inputs":[{"name":"from","type":"address"},{"name":"favor","type":"string"},{"name":"reward","type":"uint256"}],"name":"MakeRequest","outputs":[],"payable":false,"type":"function"},{"constant":true,"inputs":[{"name":"id","type":"uint256"}],"name":"GetRequest","outputs":[{"name":"from","type":"address"},{"name":"favor","type":"string"},{"name":"bound","type":"bool"},{"name":"reward","type":"uint256"}],"payable":false,"type":"function"},{"constant":true,"inputs":[{"name":"user","type":"address"},{"name":"index","type":"uint256"}],"name":"GetPromiseAt","outputs":[{"name":"id","type":"uint256"},{"name":"owner","type":"address"},{"name":"from","type":"address"},{"name":"favor","type":"string"},{"name":"offered","type":"bool"}],"payable":false,"type":"function"},{"constant":true,"inputs":[{"name":"user","type":"address"}],"name":"GetPromiseCount","outputs":[{"name":"","type":"uint256"}],"payable":false,"type":"function"},{"constant":true,"inputs":[{"name":"id","type":"uint256"}],"name":"GetPromise","outputs":[{"name":"owner","type":"address"},{"name":"from","type":"address"},{"name":"favor","type":"string"},{"name":"offered","type":"bool"}],"payable":false,"type":"function"},{"constant":true,"inputs":[{"name":"user","type":"address"}],"name":"GetRequestCount","outputs":[{"name":"","type":"uint256"}],"payable":false,"type":"function"},{"constant":false,"inputs":[{"name":"index","type":"uint256"},{"name":"id","type":"uint256"}],"name":"DropRequest","outputs":[],"payable":false,"type":"function"},{"inputs":[],"payable":false,"type":"constructor"}];

// favornet is a live contract interface into the Favor Network.
var favornet = web3.eth.contract(favornetABI).at(favornetAddress);

// address is the user's Ethereum address. We need to pull in out because it's
// contexts providing it are not accessible in suggestions...
var address;

// suggestionsContainerStyle is a simple style to make displaying suggestions
// nicer.
function suggestionsContainerStyle(count) {
  return {
    marginVertical: 1,
    marginHorizontal: 0,
    keyboardShouldPersistTaps: "always",
    height: Math.min(150, (56 * count)),
    backgroundColor: "white",
    borderRadius: 5,
    flexGrow: 1
  };
}

// The init listener simply pushes some nice messages to the user to make the
// whole experience a bit friendlier.
status.addListener("init",
  function (params, context) {
    address = context.from;

    status.sendMessage("~\"Those who want respect, give respect\"~ ~Tony Soprano");
    status.sendMessage("Never forget a favor; and never break a promise. Reputation is above all. We'll make sure of it!");
    status.sendMessage("I can list the favors you */asked* from others and those */owed* to you by others.");
});

// The asking command lists all currently open favor requests asked by the user
// from others. They may already be accepted, or just under consideration.
status.command({
  name: "asked",
  description: "Lists the favors asked from others",
  color: "#2c3e50",
  preview: function (params, context) {
    var text = status.components.text({}, "Am I asking anyone favors?");
    return {markup: status.components.view({}, [text])};
  },
  handler: function (params, context) {
    // Get the number of requests and bail if none can be found
    var requests = favornet.GetRequestCount("0x" + context.from);
    if (requests == 0) {
      status.sendMessage("You are not asking anyone favors. Good for you!");
      return;
    }
    // Yup, favors everywhere
    for (var i = 0; i < requests; i++) {
      var request = favornet.GetRequestAt("0x" + context.from, i);
      var state = "They have *not accepted* the favor request yet so you may still */drop* it."
      if (request[3]) {
        state = "They have *accepted* the favor request, you can only */honour* or */challenge* it."
      }
      var reward = " a *new promise* to return the favor";
      if (request[4] != 0) {
        var promise = favornet.GetPromise(request[4]);
        reward = "\n\n~" + promise[2] + "~\n\nfrom *" + request[1].substring(0, 8) + "…" + request[1].substring(36, 42) + "*."
      }
      status.sendMessage("Asked *" + request[1].substring(0, 8) + "…" + request[1].substring(36, 42) + "* for\n\n~" + request[2] + "~\n\nin exchange for" + reward + ".\n\n" + state);
    }
  }
});

// The owed command lists all favors currently promised by others to the user.
// The favor promises may already be offered as a reward to some user request.
status.command({
  name: "owed",
  description: "Lists the favors owed by others",
  color: "#2c3e50",
  preview: function (params, context) {
    var text = status.components.text({}, "Is anyone owing me favors?");
    return {markup: status.components.view({}, [text])};
  },
  handler: function (params, context) {
    // Get the number of requests and bail if none can be found
    var promises = favornet.GetPromiseCount("0x" + context.from);
    if (promises == 0) {
      status.sendMessage("Nope! Perhaps you should help people a bit and earn some?");
      return;
    }
    // Yup, favors everywhere
    for (var i = 0; i < promises; i++) {
      var promise = favornet.GetPromiseAt("0x" + context.from, i);
      var state = "You have *not offered* this favor to anyone, so you may */drop* it."
      if (promise[4]) {
        state = "You have *offered* this favor to someone in a request, so you cannot touch it."
      }
      status.sendMessage("Owed by *" + promise[2].substring(0, 8) + "…" + promise[2].substring(36, 42) + "*: \n\n~" + promise[3] + "~\n\n" + state);
    }
  }
});

// The requests command list all of my currently open favor requests.
status.command({
  name: "drop",
  title: "Drop",
  description: "Drops an unaccepted favor request",
  color: "#2c3e50",
  params: [{
    name: "id",
    type: status.types.TEXT,
    placeholder: "Request ID to drop"
    suggestions: dropSuggestions,
  }]
  preview: function (params, context) {
    var text = status.components.text({}, "Please drop request #" + params.id);
    return {markup: status.components.view({}, [text])};
  },
  handler: function (params, context) {
    // Find the request of the given ID and issue a transaction to drop it
    var requests = favornet.GetRequestCount("0x" + context.from);
    for (var i = 0; i < requests; i++) {
      var request = favornet.GetRequestAt("0x" + context.from, i);
      if (request[0] == params.id) {
        favornet.DropRequest.sendTransaction(i, params.id, {from: context.from}, function (error, hash) {
          if (error) {
            status.sendMessage("Favor request drop denied due to ~" + error + "~.");
          } else {
            status.sendMessage("Dropping favor request:\nhttps://ropsten.etherscan.io/tx/" + hash)
          }
        });
      }
    }
  }
});

// dropSuggestions pre-fills the suggestion box with unaccepted favor requests
// that the user may drop out of the block-chain.
function dropSuggestions() {
  // Find all the dropable favor requests
  var requests = favornet.GetRequestCount("0x" + address);

  var dropable = [];
  for (var i = 0; i < requests; i++) {
    var request = favornet.GetRequestAt("0x" + address, i);
    if (!request[3]) {
      dropable.push(request);
    }
  }
  // Render all the requests into a tapable list
  var suggestions = dropable.map(function(entry) {
    return status.components.touchable(
      {onPress: status.components.dispatch([status.events.SET_COMMAND_ARGUMENT, [0, entry[0]]])},
      status.components.view(
        suggestionsContainerStyle,
        [status.components.view(
          {borderBottomWidth: 1, borderBottomColor: "#0000001f"},
          [
            status.components.text({style: {fontWeight: "bold", marginBottom: 4}}, entry[1]),
            status.components.text({style: {fontStyle: "italic"}}, entry[2]),
          ]
        )]
      )
    );
  });
  // Give back the whole thing inside an object.
  return {markup: status.components.scrollView(suggestionsContainerStyle(suggestions.length), suggestions)};
}


// The requests command list all of my currently open favor requests.
status.command({
  name: "honour",
  title: "Honour",
  description: "Honours an accepted favor request",
  color: "#2c3e50",
  params: [{
    name: "id",
    type: status.types.TEXT,
    placeholder: "Request-reward ID to honour"
    suggestions: honourSuggestions,
  }]
  preview: function (params, context) {
    var text = status.components.text({}, "Please honour request #" + params.id);
    return {markup: status.components.view({}, [text])};
  },
  handler: function (params, context) {
    // Find the request of the given ID to allow honouring it
    var requests = favornet.GetRequestCount("0x" + context.from);
    for (var i = 0; i < requests; i++) {
      var request = favornet.GetRequestAt("0x" + context.from, i);
      if (request[0] == params.id) {
        // Request found, find the reward index if existing one
        var index = 0;
        if (request[4] != 0) {
          var promises = favornet.GetPromiseCount("0x" + context.from);
          for (var j = 0; j < requests; j++) {
            var promise = favornet.GetPromiseAt("0x" + context.from, j);
            if (promise[0] == request[4]) {
              index = j;
              break;
            }
          }
        }
        favornet.HonourRequest.sendTransaction(i, params.id, index, {from: context.from}, function (error, hash) {
          if (error) {
            status.sendMessage("Favor request honour denied due to ~" + error + "~.");
          } else {
            status.sendMessage("Honouring favor request:\nhttps://ropsten.etherscan.io/tx/" + hash)
          }
        });
        break;
      }
    }
  }
});

// honourSuggestions pre-fills the suggestion box with accepted favor requests
// that the user may honour with a promise.
function honourSuggestions() {
  // Find all the honourable favor requests
  var requests = favornet.GetRequestCount("0x" + address);

  var honourable = [];
  for (var i = 0; i < requests; i++) {
    var request = favornet.GetRequestAt("0x" + address, i);
    if (request[3]) {
      honourable.push(request);
    }
  }
  // Render all the requests into a tapable list
  var suggestions = honourable.map(function(entry) {
    return status.components.touchable(
      {onPress: status.components.dispatch([status.events.SET_COMMAND_ARGUMENT, [0, entry[0]]])},
      status.components.view(
        suggestionsContainerStyle,
        [status.components.view(
          {borderBottomWidth: 1, borderBottomColor: "#0000001f"},
          [
            status.components.text({style: {fontWeight: "bold", marginBottom: 4}}, entry[1]),
            status.components.text({style: {fontStyle: "italic"}}, entry[2]),
          ]
        )]
      )
    );
  });
  // Give back the whole thing inside an object.
  return {markup: status.components.scrollView(suggestionsContainerStyle(suggestions.length), suggestions)};
}

// The global command contains the all the different transactional interactions.
status.command({
  name: "global",
  title: "Favor Network",
  description: "Request, accept and fulfill favors",
  color: "#2c3e50",
  params: [
    {
      name: "action",
      type: status.types.TEXT,
      suggestions: globalSuggestions,
    },
    {
      name: "favor",
      type: status.types.TEXT,
      placeholder: "Favor you would like to ask"
    }
  ]
  sequentialParams: true,
  preview: function (params, context) {
    var text;
    if (params.action == "ask" || params.action.indexOf("offer-") === 0) {
      text = status.components.text({}, "Requesting favor: " + params.favor);
    } else {
      text = status.components.text({}, "Accepting favor request: " + params.favor);
    }
    return {markup: status.components.view({}, [text])};
  },
  handler: function (params, context) {
    // If we're asking for a favor, inject and return
    if (params.action == "ask" || params.action.indexOf("offer-") === 0) {
      var reward = 0;
      if (params.action.indexOf("offer-") === 0) {
        reward = params.action.substring(6);
      }
      favornet.MakeRequest.sendTransaction("0x" + context.to, params.favor, reward, {from: context.from}, function (error, hash) {
        if (error) {
          status.sendMessage("Favor request denied due to ~" + error + "~.");
        } else {
          status.sendMessage("Asked *0x" + context.to.substring(0, 6) + "…" + context.to.substring(34, 40) + "* for\n\n~" + params.favor + "~\n\nhttps://ropsten.etherscan.io/tx/" + hash)
        }
      });
      return;
    }
    // Apparently we're accepting a favor request, find it's index and bind
    var id = params.action.substring(7);

    var requests = favornet.GetRequestCount("0x" + context.to);
    for (var i = 0; i < requests; i++) {
      var request = favornet.GetRequestAt("0x" + context.to, i);
      if (request[0] == id) {
        favornet.AcceptRequest.sendTransaction("0x" + context.to, i, id, {from: context.from}, function (error, hash) {
          if (error) {
            status.sendMessage("Favor request acceptance denied due to ~" + error + "~.");
          } else {
            status.sendMessage("Accepted *0x" + context.to.substring(0, 6) + "…" + context.to.substring(34, 40) + "*:\n\n~" + request[2] + "~\n\nhttps://ropsten.etherscan.io/tx/" + hash)
          }
        });
      }
    }
  }
});

// globalSuggestions pre-fills the suggestion box with a few global actions that
// the user can do with the chat box, namely asking for a new favor or accepting
// a favor request.
function globalSuggestions(params, context) {
  // Find all the acceptable favor requests
  var requests = favornet.GetRequestCount("0x" + context.to);

  var acceptable = [];
  for (var i = 0; i < requests; i++) {
    var request = favornet.GetRequestAt("0x" + context.to, i);
    if (!request[3] && request[1] == "0x" + context.from) {
      acceptable.push(request);
    }
  }
  // Flatten the acceptable favor requests into suggestions
  var suggestions = [];

  for (var i = 0; i < acceptable.length; i++) {
    var reward = " a new promise to return the favor";
    if (request[4] != 0) {
      var promise = favornet.GetPromise(request[4]);
      reward = "\n\n~" + promise[2] + "~\n\nfrom *" + request[1].substring(0, 8) + "…" + request[1].substring(36, 42) + "*."
    }
    suggestions.push(status.components.touchable(
      {onPress: status.components.dispatch([status.events.SET_COMMAND_ARGUMENT, [0, "accept-" + acceptable[i][0]]])},
      status.components.view(
        suggestionsContainerStyle,
        [status.components.view(
          {borderBottomWidth: 1, borderBottomColor: "#0000001f"},
          [
            status.components.text({style: {marginBottom: 4}}, "Accept favor request:"),
            status.components.text({style: {marginBottom: 4, fontStyle: "italic"}}, acceptable[i][2]),
            status.components.text({style: {}}, "In exchange for: " + reward),
          ]
        )]
      )
    ));
  }
  // Suggest asking for a favor, making a new promise to repay
  suggestions.push(status.components.touchable(
    {onPress: status.components.dispatch([status.events.SET_COMMAND_ARGUMENT, [0, "ask"]])},
    status.components.view(
      suggestionsContainerStyle,
      [status.components.view(
        {borderBottomWidth: 1, borderBottomColor: "#0000001f"},
        [status.components.text({style: {}}, "Ask for a favor, promise to return it yourself")]
      )]
    )
  ));
  // Find all the giftable favor promises
  var promises = favornet.GetPromiseCount("0x" + address);

  var giftable = [];
  for (var i = 0; i < promises; i++) {
    var promise = favornet.GetPromiseAt("0x" + address, i);
    if (!promise[4]) {
      giftable.push(promise);
    }
  }
  // Flatten the giftable favor promises into suggestions
  for (var i = 0; i < giftable.length; i++) {
    suggestions.push(status.components.touchable(
      {onPress: status.components.dispatch([status.events.SET_COMMAND_ARGUMENT, [0, "offer-" + giftable[i][0]]])},
      status.components.view(
        suggestionsContainerStyle,
        [status.components.view(
          {borderBottomWidth: 1, borderBottomColor: "#0000001f"},
          [
            status.components.text({style: {marginBottom: 4}}, "Ask for a favor, reward with owned promise:"),
            status.components.text({style: {fontWeight: "bold", marginBottom: 4}}, giftable[i][2]),
            status.components.text({style: {fontStyle: "italic"}}, giftable[i][3]),
          ]
        )]
      )
    ));
  }
  // Give back the whole thing inside an object.
  return {markup: status.components.scrollView(suggestionsContainerStyle(suggestions.length), suggestions)};
}
