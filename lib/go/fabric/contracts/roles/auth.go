package roles

import (
	authpb "github.com/nova38/thesis/lib/go/gen/auth/v1"
)

func (ctx *RolesTxCtx) Authorize(ops []*authpb.Operation) (bool, error) {
	ctx.GetLogger().Info("NoAuthContract.Authenticate")

	for _, op := range ops {
		ctx.GetLogger().Info(op.String())

		// policy.ValidateOperation()
	}

	// // if the user is already authorized, return the value
	// if ctx.authChecked {
	// 	return ctx.authorized, nil
	// }
	// ctx.authChecked = true

	// if len(ops) == 0 {
	// 	return false, oops.Errorf("operations is empty")
	// }

	// ctx.ops = ops[0]

	// fn := ctx.GetAuthenticator()

	// // Check if all the items are set
	// if fn == nil {
	// 	return false, oops.Errorf("authenticator function is not set")
	// }

	// // Call the authenticator function
	// auth, err = fn(ctx, ops)

	// ctx.Logger.Info("Authorize", slog.Any("auth", auth), slog.Any("err", err))

	// if auth {
	// 	// newOps := &authpb.Operation{}

	// 	// if err = proto.Unmarshal(res.Payload, newOps); err == nil {
	// 	// 	ctx.ops.Paths = newOps.Paths
	// 	// }

	// 	return true, nil
	// }
	// return false, oops.With("operation", ops).Wrapf(err, "failed to authorize operation")

	return true, nil
}
