# apache artemis v2.54.0
console username: `artemis-activemq`

console password: `artemis-activemq`

console URL: http://localhost:8161

## URL
`amqp://localhost:5672`
## address
`tiny.profile`

## multicast (or topic)
`login.attempts`

## add user
```
/var/lib/artemis-test-cluster/bin/artemis user add \
  --user-command-user svc-profile \
  --user-command-password test \
  --role profile-producer
```
```
/var/lib/artemis-test-cluster/bin/artemis user add \
  --user-command-user svc-auth \
  --user-command-password test \
  --role profile-consumer
```

### producer: username:`svc-profile` password:`test`
### consumer: username:`svc-auth` password:`test`
### consumer: `svc-some_service` :-)

# login attempts topic settings (artemis)
```xml
         <!-- tiny profile -->
         <security-setting match="tiny.profile#">
            <permission type="manage" roles="amq"/>
            <permission type="createNonDurableQueue" roles="amq"/>
            <permission type="deleteNonDurableQueue" roles="amq"/>
            <permission type="createDurableQueue" roles="amq"/>
            <permission type="deleteDurableQueue" roles="amq"/>
            <permission type="createAddress" roles="amq"/>
            <permission type="deleteAddress" roles="amq"/>
            <permission type="consume" roles="amq,profile-consumer"/>
            <permission type="browse" roles="amq,profile-producer,profile-consumer"/>
            <permission type="send" roles="amq,profile-producer"/>
         </security-setting>
```
```xml ToDO!!!
         <!-- tiny.profile -->
         <address-setting match="tiny.profile::login.attempts">
            <dead-letter-address>tiny.auth.DLQ::login.attempts.DLQ</dead-letter-address>
            <expiry-address>tiny.auth.Expiry::login.attempts.Expiry</expiry-address>
         </address-setting>
         <address-setting match="tiny.auth::login.register">
            <dead-letter-address>tiny.auth.DLQ::login.register.DLQ</dead-letter-address>
            <expiry-address>tiny.auth.Expiry::login.register.Expiry</expiry-address>
         </address-setting>
```
```xml ToDO!!!
         <!-- tiny profile -->
         <address name="tiny.profile.DLQ">
            <anycast>
               <queue name="login.attempts.DLQ" />
               <queue name="login.register.DLQ" />
            </anycast>
         </address>
         <address name="tiny.auth.Expiry">
            <anycast>
               <queue name="login.attempts.Expiry" />
               <queue name="login.register.Expiry" />
            </anycast>
         </address>

         <address name="tiny.auth">
            <multicast>
               <queue name="login.attempts" />
               <queue name="login.register" />
            </multicast>
         </address>
```

## send test message (artemis cli)
`producer --destination topic://tiny.auth::login.attempts --message "test" --message-count 1`

## browse messages (artemis cli)
`browser --destination tiny.auth::login.attempts`

## consume message (artemis cli)
`consumer --user=auth-audit --password=test --destination topic://tiny.auth::login.attempts --message-count 1`

