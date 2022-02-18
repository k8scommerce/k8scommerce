apiServices=client\
admin

rpcServices=customer\
email\
inventory\
othersbought\
payment\
shipping\
similarproducts\
store\
user\
warehouse\
cart\
catalog


.PHONY: start
start:
	@etcd > /dev/null 2>&1 &
	@port=65000; \
	for service in $(rpcServices); do \
		printf "$(BLUE)Starting RPC Service: $(WHITE)$$service::$$port$(RESET)\n"; \
		(sleep 0.1); \
		(sed -i -e "s/:8080/:$$port/g" etc/$$service.yaml &); \
		(./$$service -f etc/$$service.yaml &); \
		echo ""; \
		port=$$((port+1)); \
	done
	@sleep .1
	@port=80; \
	for service in $(apiServices); do \
		printf "$(BLUE)Starting API Service: $(WHITE)$$service::$$port$(RESET)\n"; \
		(sleep 0.1); \
		(sed -i -e "s/: 8888/: $$port/g" etc/$$service.yaml &); \
		(./$$service -f etc/$$service.yaml &); \
		echo ""; \
		port=$$((port+8000)); \
	done

.PHONY: stop
stop:
	@for service in $(rpcServices); do \
		printf "$(BLUE)Stopping RPC Service: $(WHITE)$$service$(RESET)\n"; \
		pkill -9 -f $$service.yaml; \
		echo ""; \
	done
	@for service in $(apiServices); do \
		printf "$(BLUE)Stopping API Service: $(WHITE)$$service$(RESET)\n"; \
		pkill -9 -f $$service.yaml; \
		echo ""; \
	done
	@killall etcd
