import { server } from './server'
import express from 'express'
import { db } from './db'
import { sign } from 'jsonwebtoken'

async function startServer() {
  await server.start()
  const app = express()

  app.use(express.json())

  app.post('/login', async (req, res) => {
    const user = await db.user.findUnique({
      where: {
        email: req.body.email,
      },
    })

    if (!user) {
      res.status(401).send({ message: 'Invalid email or password' })
      return
    }

    sign(
      { uid: user.id },
      process.env.JWT_SECRET || 'super-secret',
      (err: Error | null, token?: string) => {
        if (err) {
          res.status(500).send({ message: 'Error generating token' })
          return
        }

        res.send({ token })
      }
    )
  })

  server.applyMiddleware({ app })

  await new Promise((r: any) => app.listen({ port: 4000 }, r))

  console.log('🚀 Server ready at http://localhost:4000' + server.graphqlPath)
}

startServer()
