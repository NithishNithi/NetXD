import http from 'k6/http';
import { check } from 'k6';

export default function () {
  const url = 'http://localhost:8080/api'; 
  const requestBody = {
    cusid: 1,
    firstname: "Nithish",
    lastname: "T",
    bankid: "kvb1",
    balance: 55000,
    isactive: true
  };
  const params = {
    headers: {
      'Content-Type': 'application/json',
       Authorization :"eyJhbGciOiJIUzI1NiIsInR5cCI6IkpXVCJ9.eyJjdXNpZCI6MTIzLCJleHAiOjE2OTQ1MDY0Njd9.8BOXvlDP6oJN_TvFM62pauRgZqGhhgQCHNu9tHUWU2E"
    },
  };
  const response = http.post(url, JSON.stringify(requestBody), params);
  check(response, { 'status is 200': (r) => r.status === 200 });
}
