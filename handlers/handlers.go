package handlers

import (
	"encoding/json"
	"github.com/emailtovamos/GoAPI/accounts"
	u "github.com/emailtovamos/GoAPI/utils"
	"github.com/rs/zerolog/log"
	metav1 "k8s.io/apimachinery/pkg/apis/meta/v1"
	"net/http"

	"context"
	"fmt"
	"k8s.io/apimachinery/pkg/api/errors"
	"k8s.io/client-go/kubernetes"
	"k8s.io/client-go/rest"
)

var CreateAccount = func(w http.ResponseWriter, r *http.Request) {

	account := &accounts.Account{}
	err := json.NewDecoder(r.Body).Decode(account) //decode the request body into struct and failed if any error occur
	if err != nil {
		u.Respond(w, u.Message(false, "Invalid request"))
		return
	}

	resp := account.Create() //Create accounts
	u.Respond(w, resp)
}

var Authenticate = func(w http.ResponseWriter, r *http.Request) {

	givenAccount := &accounts.Account{}
	err := json.NewDecoder(r.Body).Decode(givenAccount) //decode the request body into struct and failed if any error occur
	if err != nil {
		u.Respond(w, u.Message(false, "Invalid request"))
		return
	}

	resp := accounts.Login(givenAccount.Email, givenAccount.Password)
	u.Respond(w, resp)
}

var GetRoles = func(w http.ResponseWriter, r *http.Request) {
	input := &accounts.Input{}
	err := json.NewDecoder(r.Body).Decode(input)
	if err != nil {
		u.Respond(w, u.Message(false, "Invalid request for Getting Roles"))
		return
	}

	resp := getRoles(input)
	u.Respond(w, resp)
}

func getRoles(i *accounts.Input) map[string]interface{} {
	resp := u.Message(true, "Getting roles")
	allRoles := getRolesFromKubernetes()
	var all accounts.Roles
	for _, role := range allRoles {
		all.Roles = append(all.Roles, accounts.Role{
			Subject: "testSubject",
			Role: role,
		})
	}
	resp["roles"] = all


	//resp["accounts"] = accounts.Role{
	//	Subject: "testSubject",
	//	Role: getRolesFromKubernetes()[0],
	//} // TODO Get roles here by calling Kubernetes API

	return resp
}

func getRolesFromKubernetes() []string{
	var names []string
	// creates the in-cluster config
	config, err := rest.InClusterConfig()
	if err != nil {
		panic(err.Error())
	}
	// creates the clientset
	clientset, err := kubernetes.NewForConfig(config)
	if err != nil {
		panic(err.Error())
	}
	//for {
	// get pods in all the namespaces by omitting namespace
	// Or specify namespace to get pods in particular namespace
	pods, err := clientset.CoreV1().Pods("").List(context.TODO(), metav1.ListOptions{})
	if err != nil {
		panic(err.Error())
	}
	fmt.Printf("There are %d pods in the cluster\n", len(pods.Items))

	// Examples for error handling:
	// - Use helper functions e.g. errors.IsNotFound()
	// - And/or cast to StatusError and use its properties like e.g. ErrStatus.Message
	_, err = clientset.CoreV1().Pods("default").Get(context.TODO(), "example-xxxxx", metav1.GetOptions{})
	//roles, err := clientset.RbacV1().ClusterRoles().Get(context.TODO(), "k", metav1.GetOptions{})
	roles, err := clientset.RbacV1().ClusterRoles().List(context.TODO(),metav1.ListOptions{})
	if errors.IsNotFound(err) {
		fmt.Printf("Pod example-xxxxx not found in default namespace\n")
	} else if statusError, isStatus := err.(*errors.StatusError); isStatus {
		fmt.Printf("Error getting pod %v\n", statusError.ErrStatus.Message)
	} else if err != nil {
		//panic(err.Error())
		log.Error().Err(err).Msg("Error in getting Roles")
	} else {
		fmt.Printf("Found role in default namespace\n")
		for _, item := range roles.Items {
			fmt.Println("role name: ", item.Name)
			names = append(names, item.Name)
		}
	}

	return names
}