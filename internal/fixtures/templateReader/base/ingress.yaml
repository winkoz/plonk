apiVersion: networking.k8s.io/v1beta1
kind: Ingress
metadata:
  name: $NAME
  annotations:
    kubernetes.io/ingress.class: nginx
    cert-manager.io/cluster-issuer: letsencrypt-prod
spec:
  tls:
    - hosts:
      - $HOSTNAME
      secretName: $NAME-tls
  rules:
  - host: $HOSTNAME
    http:
      paths:
      - backend:
          serviceName: $NAME
          servicePort: $SERVICE_PORT