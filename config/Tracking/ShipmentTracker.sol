// SPDX-License-Identifier: MIT
pragma solidity ^0.8.20;

contract ShipmentTracker {
    struct Shipment {
        string shipmentId;
        address owner;
        address receiver;
        string completedStep;
        uint256 acceptedCount;
        bool exists;
    }

    mapping(string => Shipment) public shipments;

    event ShipmentCreated(
        string shipmentId,
        address indexed owner,
        address indexed receiver
    );

    function createShipment(
        string memory _shipmentId,
        address _receiver,
        string memory _completedStep,
        uint256 _acceptedCount
    ) public {
        require(!shipments[_shipmentId].exists, "Shipment already exists");

        shipments[_shipmentId] = Shipment({
            shipmentId: _shipmentId,
            owner: msg.sender,
            receiver: _receiver,
            completedStep: _completedStep,
            acceptedCount: _acceptedCount,
            exists: true
        });

        emit ShipmentCreated(_shipmentId, msg.sender, _receiver);
    }

    function getShipment(string memory _shipmentId)
        public
        view
        returns (
            string memory,
            address,
            address,
            string memory,
            uint256
        )
    {
        require(shipments[_shipmentId].exists, "Shipment not found");
        Shipment memory s = shipments[_shipmentId];
        return (
            s.shipmentId,
            s.owner,
            s.receiver,
            s.completedStep,
            s.acceptedCount
        );
    }
}
