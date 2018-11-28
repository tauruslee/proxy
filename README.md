curl -X GET http://host:port/proxy/info
	select all queueId
	{
		responseTime: [response time from submit to response],
		status: success or failed reason,
		result: [
			{queueId:"queueId1"},
			{queueId:"queueId2"},
			{queueId:"queueId3"},
		],
	}

curl -X GET http://host:port/proxy/info/{queueId}
	select all key for queueId
	{
		responseTime: response time from submit to response,
		status: success or failed reason,
		result: [
			{id:"id1"},
			{id:"id2"},
			{id:"id3"},
		],
	}

curl -X DELETE http://host:port/proxy/{queueId}
	delete queueId
	{
		responseTime: response time from submit to response,
		status: success or failed reason,
		result: deleted queueId,
	}

curl -X POST http://host:port/proxy/{queueId}
	insert queueId
	{
		responseTime: response time from submit to response,
		status: success or failed reason,
		result: id,
	}


curl -X GET http://host:port/proxy/{queueId}/{id}
	select queueId record id
	{
		responseTime: response time from submit to response,
		status: success or failed reason,
		result: record retrieved,
	}

curl -X POST -d info.json http://host:port/proxy/{queueId}
	insert to queueId
	{
		responseTime: response time from submit to response,
		status: success or failed reason,
		result: id,
	}

curl -X DELETE http://host:port/proxy/{queueId}/{id}
	delete queueId record id
	{
		responseTime: response time from submit to response,
		status: success or failed reason,
		result: id,
	}

