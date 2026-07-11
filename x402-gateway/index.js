import express from 'express';
import { createProxyMiddleware } from 'http-proxy-middleware';
import { injectivePaymentMiddleware } from '@injectivelabs/x402/middleware';

const app = express();
const BACKEND_URL = process.env.BACKEND_URL || 'http://localhost:8080';
const FACILITATOR_URL = process.env.X402_FACILITATOR_URL || 'https://x402.org/facilitator';
const GATEWAY_PORT = process.env.GATEWAY_PORT || 3000;

// x402 payment config: 0.01 USDC on Injective for premium predictions.
const paymentConfig = {
  'GET /api/premium-predict/:id': {
    accepts: [{
      network: 'eip155:1776',
      asset: '0xa00C59fF5a080D2b954d0c75e46E22a0c371235a',
      amount: '10000', // 0.01 USDC (6 decimals)
    }],
  },
};

app.use(injectivePaymentMiddleware(paymentConfig, { facilitatorUrl: FACILITATOR_URL }));

// Everything else proxies to the Go backend unchanged.
app.use('/api', createProxyMiddleware({
  target: BACKEND_URL,
  changeOrigin: true,
}));

app.listen(GATEWAY_PORT, () => {
  console.log(`x402 gateway listening on http://localhost:${GATEWAY_PORT}`);
  console.log(`Premium predictions gated at 0.01 USDC; backend proxy -> ${BACKEND_URL}`);
});
