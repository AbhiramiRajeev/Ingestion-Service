# Ingestion-Service
API gateway and event collector for the Automated Incident Response Orchestrator. It securely ingests real-time security events (such as login attempts, suspicious activities, brute force signals) from various endpoints or clients.



<!-- 

curl -X POST http://localhost:8080/api/v1/events \
  -H "Authorization: Bearer mysecretkey" \
  -H "Content-Type: application/json" \
  -d '{
    "event_type": "login_attempt",
    "user_id": "alice",
    "ip_address": "1.2.3.4",
    "status": "failure",
    "timestamp": "2025-06-16T12:00:00Z"
  }' 
  
  -->
