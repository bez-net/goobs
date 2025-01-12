package e2e

import (
	"testing"

	"github.com/andreykaipov/goobs"
	"github.com/andreykaipov/goobs/api/requests/scenes"
	"github.com/andreykaipov/goobs/api/requests/sources"
	"github.com/andreykaipov/goobs/api/typedefs"
	"github.com/stretchr/testify/assert"
)

func check(err error) {
	if err != nil {
		panic(err)
	}
}

func Test_goobs_e2e(t *testing.T) {
	client, err := goobs.New("localhost:4545", goobs.WithPassword("hello"))
	assert.NoError(t, err)

	sceneName := "goobs test scene"

	t.Run("Scenes", func(t *testing.T) {
		resp, err := client.Scenes.GetSceneList()
		assert.NoError(t, err)
		list := resp.Scenes

		for _, s := range list {
			if s.Name == sceneName {
				// Prior OBS instance probably hasn't been closed
				t.Errorf("Scene list already contains e2e scene %q.", sceneName)
			}
		}

		_, err = client.Scenes.CreateScene(&scenes.CreateSceneParams{SceneName: sceneName})
		assert.NoError(t, err)
		_, err = client.Scenes.CreateScene(&scenes.CreateSceneParams{SceneName: sceneName})
		assert.Error(t, err)

		resp, err = client.Scenes.GetSceneList()
		assert.NoError(t, err)
		newList := resp.Scenes
		assert.Len(t, newList, len(list)+1)
	})

	t.Run("Sources", func(t *testing.T) {
		sourceName := "goobs test source (text)"

		resp, err := client.Sources.GetSourcesList()
		assert.NoError(t, err)
		list := resp.Sources

		for _, s := range list {
			if s.Name == sourceName {
				// Prior OBS instance probably hasn't been closed
				t.Errorf("Source list already contains e2e source %q.", sourceName)
			}
		}

		_, err = client.Sources.CreateSource(&sources.CreateSourceParams{
			SourceName: sourceName,
			SourceKind: "text_ft2_source_v2",
			SceneName:  sceneName,
			SetVisible: false,
		})
		assert.NoError(t, err)
		_, err = client.Sources.CreateSource(&sources.CreateSourceParams{
			SourceName: sourceName,
			SourceKind: "text_ft2_source_v2",
			SceneName:  sceneName,
			SetVisible: false,
		})
		assert.Error(t, err)

		resp, err = client.Sources.GetSourcesList()
		assert.NoError(t, err)
		newList := resp.Sources
		assert.Len(t, newList, len(list)+1)

		t.Run("Text", func(t *testing.T) {
			resp1, err := client.Sources.GetTextFreetype2Properties(&sources.GetTextFreetype2PropertiesParams{Source: sourceName})
			assert.NoError(t, err)
			assert.Nil(t, resp1.Font)
			assert.Empty(t, resp1.Text)

			resp2, err := client.Sources.SetTextFreetype2Properties(&sources.SetTextFreetype2PropertiesParams{
				Source: sourceName,
				Text:   "dicky long neck",
				Color1: 123,
				Font: &typedefs.Font{
					Face: "Arial",
					Size: 11,
				},
			})
			assert.NotNil(t, resp2)
			assert.NoError(t, err)

			resp3, err := client.Sources.GetTextFreetype2Properties(&sources.GetTextFreetype2PropertiesParams{Source: sourceName})
			assert.NoError(t, err)
			assert.Equal(t, resp3.Text, "dicky long neck")
			// TODO file a bug upstream
			// seems like Font is not returned in the response
			// assert.NotNil(t, resp3.Font)
		})
	})
}
