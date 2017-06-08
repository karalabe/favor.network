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

// The init listener simply pushes some nice messages to the user to make the
// whole experience a bit friendlier.
status.addListener("init",
  function (params, context) {
    address = context.from;

    status.sendMessage("~\"Those who want respect, give respect\"~ ~Tony Soprano");
    status.sendMessage("Never forget a favor; and never break a promise. Reputation is above all. We'll make sure of it!");
});

// The requests command list all of my currently open favor requests.
status.command({
  name: "requests",
  title: "Requests",
  description: "Lists my open favor requests",
  color: "#2c3e50",
  preview: function (params, context) {
    var text = status.components.text({}, "Am I asking anyone for favors?");
    return {markup: status.components.view({}, [text])};
  },
  handler: function (params, context) {
    // Get the number of requests and bail if none can be found
    var requests = favornet.GetRequestCount("0x" + context.from);
    if (requests == 0) {
      status.sendMessage("You're not asking anyone for favors. Good for you!");
      return;
    }
    // Yup, favors everywhere
    for (var i = 0; i < requests; i++) {
      var request = favornet.GetRequestAt("0x" + context.from, i);
      var state = "They have not */accept*-ed the favor request yet so you can still */drop* it."
      if (request[3]) {
        state = "Accepted, honor or challenge it."
      }
      status.sendMessage("Asking *" + request[1].substring(0, 8) + "…" + request[1].substring(36, 42) + "* for\n\n~" + request[2] + "~\n\n" + state);
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

function suggestionsContainerStyle(suggestionsCount) {
    return {
        marginVertical: 1,
        marginHorizontal: 0,
        keyboardShouldPersistTaps: "always",
        height: Math.min(150, (56 * suggestionsCount)),
        backgroundColor: "white",
        borderRadius: 5,
        flexGrow: 1
    };
}

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


// The favors command lists all of the favors currently promised to me.
status.command({
  name: "favors",
  title: "Favors",
  description: "List the favors promised to me",
  color: "#2c3e50",
  preview: function (params, context) {
    var text = status.components.text({}, "Is anyone owing me favors?");
    return {markup: status.components.view({}, [text])};
  },
  handler: function (params, context) {
    // Get the number of requests and bail if none can be found
    var promises = favornet.GetPromiseCount("0x" + context.from);
    if (promises == 0) {
      status.sendMessage("Not that I know of!");
      return;
    }
    // Yup, favors everywhere
    status.sendMessage("Yup");
    for (var i = 0; i < promises; i++) {
      var promise = favornet.GetPromiseAt("0x" + context.from, i);
      status.sendMessage(promise);
    }
  }
});

// The global command contains the all the different transactional interactions.
status.command({
  name: "global",
  title: "Favor Network",
  description: "Request, accept and fulfill favors",
  color: "#2c3e50",
  params: [{
    name: "favor",
    type: status.types.TEXT,
    placeholder: "Favor you'd like to ask"
  }]
  sequentialParams: true,
  preview: function (params, context) {
    var text = status.components.text({}, "Requesting favor: " + params.favor);
    return {markup: status.components.view({}, [text])};
  },
  handler: function (params, context) {
    favornet.MakeRequest.sendTransaction("0x" + context.to, params.favor, "0x0", {from: context.from}, function (error, hash) {
      if (error) {
        status.sendMessage("Favor request denied due to ~" + error + "~.");
      } else {
        status.sendMessage("Asked *0x" + context.to.substring(0, 6) + "…" + context.to.substring(34, 40) + "* for\n\n~" + params.favor + "~\n\nhttps://ropsten.etherscan.io/tx/" + hash)
      }
    });
  },
});
