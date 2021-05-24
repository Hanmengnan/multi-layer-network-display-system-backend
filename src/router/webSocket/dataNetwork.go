package webSocket

import (
	database "multi-layer-network-display-system-backend/src/model"
)

type DataNetworkLinkDetailResponse struct {
	Response         int                        `json:"response"`
	BasicInfo        database.LinkInfo          `json:"basicInfo"`
	ParameterChange  []database.ParameterChange `json:"parameterChange"`
	OriginStatistics map[string]int32           `json:"originStatistics"`
}
type DataNetworkNodeDetailResponse struct {
	Response         int `json:"response"`
	NodeDetail       database.NodeDetail
	LinkStatistics   []database.MessageLinkStatistics
	DateStatistics   []database.MessageDateStatistics
	OriginStatistics map[string]int32
}

func DataNetworkLinkDetail(linkId string) DataNetworkLinkDetailResponse {
	var res DataNetworkLinkDetailResponse
	res.Response = 0
	res.BasicInfo = database.GetLinkBasicInfo(linkId)
	res.ParameterChange = database.GetLinkParameterChange(linkId)
	res.OriginStatistics = database.GetLinkMessageOriginStatistics(linkId)
	return res
}

func DataNetworkNodeDetail(nodeId string) DataNetworkNodeDetailResponse {
	var res DataNetworkNodeDetailResponse
	res.Response = 0
	res.LinkStatistics = database.GetNodeMessageLinkStatistics(nodeId)
	res.DateStatistics = database.GetNodeMessageDateStatistics(nodeId)
	res.OriginStatistics = database.GetNodeMessageOriginStatistics(nodeId)
	res.NodeDetail = database.GetNodeDetail(nodeId)
	return res
}
