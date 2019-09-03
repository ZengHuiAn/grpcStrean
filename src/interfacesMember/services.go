package interfacesMember

type ServiceBehavior interface {
	Start()
	Update()
	OnDestroy()
}
