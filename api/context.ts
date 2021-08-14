import { PrismaClient, User } from '@prisma/client'
import { AuthenticationError } from 'apollo-server-express'
import { verify } from 'jsonwebtoken'
import { db } from './db'

export interface ContextInfo {
  db: PrismaClient
  user: User | null
}

export const context = async ({ req }: { req: any }): Promise<ContextInfo> => {
  const auth = req?.headers?.authorization

  if (!auth) {
    return { db, user: null }
  }

  const token = auth.split(' ')[1]

  const decoded: any = verify(token, process.env.JWT_SECRET || 'super-secret')

  if (!decoded) {
    throw new AuthenticationError('Invalid token')
  }

  const user = await db.user.findUnique({
    where: {
      id: Number(decoded.uid),
    },
  })

  return { db, user }
}
