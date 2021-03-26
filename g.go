const destinyDirectory = "/destiny/path/"
const defaultContainerName = "test-container"

func SaveDump(c *gin.Context) {
	var message BigQueueMessageForDump
	var response *http.Response

	if err := c.ShouldBind(&message); err != nil {
		c.JSON(http.StatusInternalServerError, err.Error())
		return
	}

	containerName := c.Query("c")
	if containerName == "" {
		containerName = defaultContainerName
	}

	objectStorage.ContainerName = containerName
	objectStorage.CreateOsClient()

	req, err := http.NewRequestWithContext(c, http.MethodGet, message.URL, nil)

	if err != nil {
		c.JSON(http.StatusInternalServerError, err.Error())
		return
	}

	response, err = http.DefaultClient.Do(req)
	if err != nil {
		c.JSON(http.StatusInternalServerError, err.Error())
		return
	}

	defer response.Body.Close()

	if response.StatusCode == http.StatusOK {
		bodyBytes, err := ioutil.ReadAll(response.Body)
		if err != nil {
			c.JSON(http.StatusInternalServerError, err.Error())
			return
		}

		path := message.FilePath
		fileName := filepath.Base(path)

		err = objectStorage.Put(destinyDirectory+fileName, bodyBytes)
		if err != nil && err.Error() != "" {
			c.JSON(http.StatusInternalServerError, err.Error())
			return
		}
	}

	c.JSON(http.StatusOK, "Files Uploaded!")
}