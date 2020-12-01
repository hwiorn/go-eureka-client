package eureka

import (
	"net/http"
	"strings"
)

func (c *Client) SendHeartbeat(appId, instanceId string) error {
	values := []string{"apps", appId, instanceId}
	path := strings.Join(values, "/")
	resp, err := c.Put(path, nil)
	if err != nil {
		return err
	}
	switch resp.StatusCode {
	case http.StatusNotFound:
		return newError(ErrCodeInstanceNotFound,
			"Instance resource not found when sending heartbeat", 0)
	}
	return nil
}

func (c *Client) UpdateStatus(instance *InstanceInfo, status string) error {
	values := []string{"apps", instance.App, instance.HostName, "status"}
	path := strings.Join(values, "/") + "?value=" + status
	resp, err := c.Put(path, nil)
	if err != nil {
		return err
	}
	switch resp.StatusCode {
	case http.StatusNotFound:
		return newError(ErrCodeInstanceNotFound,
			"Instance resource not found when sending status", 0)
	}
	instance.Status = status
	return nil
}
