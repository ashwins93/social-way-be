import { extendType, objectType } from 'nexus'

export const Post = objectType({
  name: 'Post',
  definition(t) {
    t.nonNull.int('id')
    t.nonNull.date('createdAt')
    t.nonNull.string('content')
  },
})

export const PostQuery = extendType({
  type: 'Query',
  definition(t) {
    t.nonNull.list.field('posts', {
      type: 'Post',
      resolve: (parent, args, ctx) => {
        return ctx.db.post.findMany()
      },
    })
  },
})
