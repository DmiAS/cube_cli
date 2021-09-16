package iproto

import (
	"github.com/DmiAS/cube_cli/internal/app/connection"
	"github.com/DmiAS/cube_cli/internal/app/models"
)

// полученный поток байт перевести в структуру пакета ответа от сервера, определить какой
// получили ответ и вернуть соответствующую структуру
func getResp(conn connection.Connection) (Response, error) {
	data, err := conn.Read()
	if err != nil {
		return nil, err
	}

	resPacket := new(models.ResponsePacket)
	if err := UnMarshal(data, resPacket); err != nil {
		return nil, err
	}

	resp, err := determineResponse(resPacket.Body)
	if err != nil {
		return nil, err
	}
	return resp, nil
}

// следует определить по код какую структуру сформировать
func determineResponse(packet models.ResponseBody) (Response, error) {
	if packet.ReturnCode != 0 {
		errString, err := packet.ErrorString.ToString()
		if err != nil {
			return nil, err
		}
		return models.ResponseErr{
			ReturnCode:  packet.ReturnCode,
			ErrorString: errString,
		}, nil
	}

	clientID, err := packet.ClientID.ToString()
	if err != nil {
		return nil, err
	}

	userName, err := packet.UserName.ToString()
	if err != nil {
		return nil, err
	}
	return models.ResponseOk{
		ReturnCode: packet.ReturnCode,
		ClientID:   clientID,
		ClientType: packet.ClientType,
		UserName:   userName,
		ExpiresIn:  packet.ExpiresIn,
		UserID:     packet.UserID,
	}, nil
}
