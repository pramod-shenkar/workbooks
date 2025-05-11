// SPDX-License-Identifier: UNKNOWN 
pragma solidity ^0.8.29;

library requestmodel {
    enum RequestStatus {
        Created,
        Approved,
        Rejected,
        Failed,
        Processed
    }

    struct Request {
        uint256 storageId;
        string hashdata;
        address creator;
        RequestStatus status;
    }
}

contract requeststore {
    mapping(uint256 => requestmodel.Request) private requests;

    function SaveRequest(
        uint256 id,
        requestmodel.Request memory _request
    ) public {
        _request.status = requestmodel.RequestStatus.Created;
        requests[id] = _request;
    }

    function GetRequest(
        uint256 id
    ) public view returns (requestmodel.Request memory) {
        return requests[id];
    }
}

contract RequestContract {
    event SavedEvent(uint256 indexed id, bool status);
    event ApprovedEvent(uint256 indexed id, bool status);

    requeststore store;

    function saveRequest(
        uint256 id,
        requestmodel.Request memory _request
    ) public {
        _request.status = requestmodel.RequestStatus.Created;
        store.SaveRequest(id, _request);
        emit SavedEvent(id, true);
    }

    function approveRequest(uint256 id) public {
        requestmodel.Request memory request = store.GetRequest(id);
        require(id != 0, "Request not found");
        request.status = requestmodel.RequestStatus.Approved;
        store.SaveRequest(id, request);
        emit ApprovedEvent(id, true);
    }

    function queryRequest(
        uint256 id
    ) public view returns (requestmodel.Request memory) {
        return store.GetRequest(id);
    }
}
