package delivery

import (
	"be13/project/features/feedback"
)

type FeedbackRequest struct {
	Status   string `json:"status" form:"status"`
	MenteeID uint   `json:"mentee_id" form:"mentee_id"`
	Comment  string `json:"comment" form:"comment"`
}

// func (req *FeedbackRequest) reqToCore() feedback.Core {
// 	return feedback.Core{
// 		// ClassName: req.ClassName,
// 		UserID:   req.UserID,
// 		Status:   req.Status,
// 		MenteeID: req.MenteeID,
// 		Comment:  req.Comment,
// 	}

// }

func FeedbackRequestToUserCore(data FeedbackRequest) feedback.Core {
	return feedback.Core{

		Status:   data.Status,
		MenteeID: data.MenteeID,
		Comment:  data.Comment,
	}
}
