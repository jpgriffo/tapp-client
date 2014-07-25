package script

import "encoding/json"
import "os"
import "os/exec"
import "syscall"
import "time"

type ScriptCharacterization struct {
	Type string `json:"type"`
	ExecutionOrder int64 `json:"execution_order"`
	Uuid string `json:"uuid"`
	Script Script `json:"script"` 
}

type Script struct {
	Uuid string `json:"uuid"`
	Code string `json:"code"`
	AttachmentPaths []string `json:"attachment_paths"`
}

type ScriptConclusion struct {
	ScriptCharacterizationId string `json:"script_characterization_id"`
	Output string `json:"output"`
	ExitCode int `json:"exit_code"`
	StartedAt time.Time `json:"started_at"`
	FinishedAt time.Time `json:"finished_at"`
}

func (script_characterization ScriptCharacterization) ToFile() (string, error){
	file_dir_path := "/tmp/tappscript/"
	file_path := file_dir_path + script_characterization.Uuid + ".sh"
	err := os.MkdirAll(file_dir_path, 0777)
	if (err == nil) {
		file, err := os.Create(file_path)
		if (err == nil){
			defer file.Close()
			_, err := file.WriteString(script_characterization.Script.Code)
			if (err == nil){
				file.Sync()
			}
		}
	}
	return file_path, err
}

func (script_characterization ScriptCharacterization) Execute() (*ScriptConclusion, error) {
	var exit_code int
	var output []byte
	var startTime time.Time
	var endTime time.Time
	script_filepath, err := script_characterization.ToFile()
	if (err == nil) {
		cmd := exec.Command("sh", script_filepath)
		startTime = time.Now()
		output, err = cmd.Output()
		endTime = time.Now()
		if (err != nil){ 
			if exiterr, ok := err.(*exec.ExitError); ok {
				if status, ok := exiterr.Sys().(syscall.WaitStatus); ok {
					exit_code = status.ExitStatus()
				}
			}
		}
	}
	return &ScriptConclusion{script_characterization.Uuid, string(output[:]), exit_code, startTime, endTime}, err
}

func (script_conclusion ScriptConclusion) ToJson() ([]byte, error) {
	return json.Marshal(script_conclusion) 
}