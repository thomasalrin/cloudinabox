/*
** Copyright [2012-2014] [Megam Systems]
**
** Licensed under the Apache License, Version 2.0 (the "License");
** you may not use this file except in compliance with the License.
** You may obtain a copy of the License at
**
** http://www.apache.org/licenses/LICENSE-2.0
**
** Unless required by applicable law or agreed to in writing, software
** distributed under the License is distributed on an "AS IS" BASIS,
** WITHOUT WARRANTIES OR CONDITIONS OF ANY KIND, either express or implied.
** See the License for the specific language governing permissions and
** limitations under the License.
 */

package servers

import (
	"github.com/megamsys/cloudinabox/app"
	"github.com/megamsys/cloudinabox/models/orm"
   "fmt"
   "errors"
   "net/http"
)


func InstallServers(serverName string) error {
	var err error
    switch serverName  {
    	case "MEGAM":
    	  err = app.MegamInstall()
    	  if err != nil {
		     fmt.Printf("Error: Install error for [%s]", serverName)
		     fmt.Println(err)
		     return err
	      }
        case "COBBLER":
          err = app.CobblerInstall()  
          if err != nil {
		     fmt.Printf("Error: Install error for [%s]", serverName)
		     fmt.Println(err)
		     return err
	      }
       case "OPENNEBULA":
          err = app.NebulaInstall()  
          if err != nil {
		     fmt.Printf("Error: Install error for [%s]", serverName)
		     fmt.Println(err)
		     return err
	      }
       case "OPENNEBULAHOST":
          err = app.OpenNebulaHostInstall()  
          if err != nil {
		     fmt.Printf("Error: Install error for [%s]", serverName)
		     fmt.Println(err)
		     return err
	      }   
    }
    
	return nil
}

func InstallNode(server *orm.Servers) error {
	url := "http://"+server.IP+":8077/servernodes/nodes/"+server.Name
    res, err := http.Get(url)
	if err != nil {
		return err
	} else {
		if res.StatusCode > 299 {
			return errors.New(res.Status)
		} else {
	        err = app.SCPSSHInstall()  
	        return err
	   }
	}
}


