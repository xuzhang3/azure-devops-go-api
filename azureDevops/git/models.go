﻿// --------------------------------------------------------------------------------------------
// Copyright (c) Microsoft Corporation. All rights reserved.
// Licensed under the MIT License.
// --------------------------------------------------------------------------------------------
// Generated file, DO NOT EDIT
// Changes may cause incorrect behavior and will be lost if the code is regenerated.
// --------------------------------------------------------------------------------------------

package git

import (
    "github.com/google/uuid"
    "time"
)

type AssociatedWorkItem struct {
    AssignedTo *string `json:"assignedTo,omitempty"`
    // Id of associated the work item.
    Id *int `json:"id,omitempty"`
    State *string `json:"state,omitempty"`
    Title *string `json:"title,omitempty"`
    // REST Url of the work item.
    Url *string `json:"url,omitempty"`
    WebUrl *string `json:"webUrl,omitempty"`
    WorkItemType *string `json:"workItemType,omitempty"`
}

type AsyncGitOperationNotification struct {
    OperationId *int `json:"operationId,omitempty"`
}

type AsyncRefOperationCommitLevelEventNotification struct {
    OperationId *int `json:"operationId,omitempty"`
    CommitId *string `json:"commitId,omitempty"`
}

type AsyncRefOperationCompletedNotification struct {
    OperationId *int `json:"operationId,omitempty"`
    NewRefName *string `json:"newRefName,omitempty"`
}

type AsyncRefOperationConflictNotification struct {
    OperationId *int `json:"operationId,omitempty"`
    CommitId *string `json:"commitId,omitempty"`
}

type AsyncRefOperationGeneralFailureNotification struct {
    OperationId *int `json:"operationId,omitempty"`
}

type AsyncRefOperationProgressNotification struct {
    OperationId *int `json:"operationId,omitempty"`
    CommitId *string `json:"commitId,omitempty"`
    Progress *float64 `json:"progress,omitempty"`
}

type AsyncRefOperationTimeoutNotification struct {
    OperationId *int `json:"operationId,omitempty"`
}

// Meta data for a file attached to an artifact.
type Attachment struct {
    // Links to other related objects.
    Links *ReferenceLinks `json:"_links,omitempty"`
    // The person that uploaded this attachment.
    Author *IdentityRef `json:"author,omitempty"`
    // Content hash of on-disk representation of file content. Its calculated by the server by using SHA1 hash function.
    ContentHash *string `json:"contentHash,omitempty"`
    // The time the attachment was uploaded.
    CreatedDate *time.Time `json:"createdDate,omitempty"`
    // The description of the attachment.
    Description *string `json:"description,omitempty"`
    // The display name of the attachment. Can't be null or empty.
    DisplayName *string `json:"displayName,omitempty"`
    // Id of the attachment.
    Id *int `json:"id,omitempty"`
    // Extended properties.
    Properties *interface{} `json:"properties,omitempty"`
    // The url to download the content of the attachment.
    Url *string `json:"url,omitempty"`
}

// Real time event (SignalR) for an auto-complete update on a pull request
type AutoCompleteUpdatedEvent struct {
    // The id of this event. Can be used to track send/receive state between client and server.
    EventId *uuid.UUID `json:"eventId,omitempty"`
    // The id of the pull request this event was generated for.
    PullRequestId *int `json:"pullRequestId,omitempty"`
}

// Real time event (SignalR) for a source/target branch update on a pull request
type BranchUpdatedEvent struct {
    // The id of this event. Can be used to track send/receive state between client and server.
    EventId *uuid.UUID `json:"eventId,omitempty"`
    // The id of the pull request this event was generated for.
    PullRequestId *int `json:"pullRequestId,omitempty"`
    // If true, the source branch of the pull request was updated
    IsSourceUpdate *bool `json:"isSourceUpdate,omitempty"`
}

type Change struct {
    // The type of change that was made to the item.
    ChangeType *VersionControlChangeType `json:"changeType,omitempty"`
    // Current version.
    Item *interface{} `json:"item,omitempty"`
    // Content of the item after the change.
    NewContent *ItemContent `json:"newContent,omitempty"`
    // Path of the item on the server.
    SourceServerItem *string `json:"sourceServerItem,omitempty"`
    // URL to retrieve the item.
    Url *string `json:"url,omitempty"`
}

type ChangeCountDictionary struct {
}

type ChangeList struct {
    AllChangesIncluded *bool `json:"allChangesIncluded,omitempty"`
    ChangeCounts *map[VersionControlChangeType]int `json:"changeCounts,omitempty"`
    Changes *[]Change `json:"changes,omitempty"`
    Comment *string `json:"comment,omitempty"`
    CommentTruncated *bool `json:"commentTruncated,omitempty"`
    CreationDate *time.Time `json:"creationDate,omitempty"`
    Notes *[]CheckinNote `json:"notes,omitempty"`
    Owner *string `json:"owner,omitempty"`
    OwnerDisplayName *string `json:"ownerDisplayName,omitempty"`
    OwnerId *uuid.UUID `json:"ownerId,omitempty"`
    SortDate *time.Time `json:"sortDate,omitempty"`
    Version *string `json:"version,omitempty"`
}

// Criteria used in a search for change lists
type ChangeListSearchCriteria struct {
    // If provided, a version descriptor to compare against base
    CompareVersion *string `json:"compareVersion,omitempty"`
    // If true, don't include delete history entries
    ExcludeDeletes *bool `json:"excludeDeletes,omitempty"`
    // Whether or not to follow renames for the given item being queried
    FollowRenames *bool `json:"followRenames,omitempty"`
    // If provided, only include history entries created after this date (string)
    FromDate *string `json:"fromDate,omitempty"`
    // If provided, a version descriptor for the earliest change list to include
    FromVersion *string `json:"fromVersion,omitempty"`
    // Path of item to search under. If the itemPaths memebr is used then it will take precedence over this.
    ItemPath *string `json:"itemPath,omitempty"`
    // List of item paths to search under. If this member is used then itemPath will be ignored.
    ItemPaths *[]string `json:"itemPaths,omitempty"`
    // Version of the items to search
    ItemVersion *string `json:"itemVersion,omitempty"`
    // Number of results to skip (used when clicking more...)
    Skip *int `json:"skip,omitempty"`
    // If provided, only include history entries created before this date (string)
    ToDate *string `json:"toDate,omitempty"`
    // If provided, the maximum number of history entries to return
    Top *int `json:"top,omitempty"`
    // If provided, a version descriptor for the latest change list to include
    ToVersion *string `json:"toVersion,omitempty"`
    // Alias or display name of user who made the changes
    User *string `json:"user,omitempty"`
}

type CheckinNote struct {
    Name *string `json:"name,omitempty"`
    Value *string `json:"value,omitempty"`
}

// Represents a comment which is one of potentially many in a comment thread.
type Comment struct {
    // Links to other related objects.
    Links *ReferenceLinks `json:"_links,omitempty"`
    // The author of the comment.
    Author *IdentityRef `json:"author,omitempty"`
    // The comment type at the time of creation.
    CommentType *CommentType `json:"commentType,omitempty"`
    // The comment content.
    Content *string `json:"content,omitempty"`
    // The comment ID. IDs start at 1 and are unique to a pull request.
    Id *int `json:"id,omitempty"`
    // Whether or not this comment was soft-deleted.
    IsDeleted *bool `json:"isDeleted,omitempty"`
    // The date the comment's content was last updated.
    LastContentUpdatedDate *time.Time `json:"lastContentUpdatedDate,omitempty"`
    // The date the comment was last updated.
    LastUpdatedDate *time.Time `json:"lastUpdatedDate,omitempty"`
    // The ID of the parent comment. This is used for replies.
    ParentCommentId *int `json:"parentCommentId,omitempty"`
    // The date the comment was first published.
    PublishedDate *time.Time `json:"publishedDate,omitempty"`
    // A list of the users who have liked this comment.
    UsersLiked *[]IdentityRef `json:"usersLiked,omitempty"`
}

// Comment iteration context is used to identify which diff was being viewed when the thread was created.
type CommentIterationContext struct {
    // The iteration of the file on the left side of the diff when the thread was created. If this value is equal to SecondComparingIteration, then this version is the common commit between the source and target branches of the pull request.
    FirstComparingIteration *int `json:"firstComparingIteration,omitempty"`
    // The iteration of the file on the right side of the diff when the thread was created.
    SecondComparingIteration *int `json:"secondComparingIteration,omitempty"`
}

type CommentPosition struct {
    // The line number of a thread's position. Starts at 1.
    Line *int `json:"line,omitempty"`
    // The character offset of a thread's position inside of a line. Starts at 0.
    Offset *int `json:"offset,omitempty"`
}

// Represents a comment thread of a pull request. A thread contains meta data about the file it was left on along with one or more comments (an initial comment and the subsequent replies).
type CommentThread struct {
    // Links to other related objects.
    Links *ReferenceLinks `json:"_links,omitempty"`
    // A list of the comments.
    Comments *[]Comment `json:"comments,omitempty"`
    // The comment thread id.
    Id *int `json:"id,omitempty"`
    // Set of identities related to this thread
    Identities *map[string]IdentityRef `json:"identities,omitempty"`
    // Specify if the thread is deleted which happens when all comments are deleted.
    IsDeleted *bool `json:"isDeleted,omitempty"`
    // The time this thread was last updated.
    LastUpdatedDate *time.Time `json:"lastUpdatedDate,omitempty"`
    // Optional properties associated with the thread as a collection of key-value pairs.
    Properties *interface{} `json:"properties,omitempty"`
    // The time this thread was published.
    PublishedDate *time.Time `json:"publishedDate,omitempty"`
    // The status of the comment thread.
    Status *CommentThreadStatus `json:"status,omitempty"`
    // Specify thread context such as position in left/right file.
    ThreadContext *CommentThreadContext `json:"threadContext,omitempty"`
}

type CommentThreadContext struct {
    // File path relative to the root of the repository. It's up to the client to use any path format.
    FilePath *string `json:"filePath,omitempty"`
    // Position of last character of the thread's span in left file.
    LeftFileEnd *CommentPosition `json:"leftFileEnd,omitempty"`
    // Position of first character of the thread's span in left file.
    LeftFileStart *CommentPosition `json:"leftFileStart,omitempty"`
    // Position of last character of the thread's span in right file.
    RightFileEnd *CommentPosition `json:"rightFileEnd,omitempty"`
    // Position of first character of the thread's span in right file.
    RightFileStart *CommentPosition `json:"rightFileStart,omitempty"`
}

// The status of a comment thread.
type CommentThreadStatus string

// Comment tracking criteria is used to identify which iteration context the thread has been tracked to (if any) along with some detail about the original position and filename.
type CommentTrackingCriteria struct {
    // The iteration of the file on the left side of the diff that the thread will be tracked to. Threads were tracked if this is greater than 0.
    FirstComparingIteration *int `json:"firstComparingIteration,omitempty"`
    // Original filepath the thread was created on before tracking. This will be different than the current thread filepath if the file in question was renamed in a later iteration.
    OrigFilePath *string `json:"origFilePath,omitempty"`
    // Original position of last character of the thread's span in left file.
    OrigLeftFileEnd *CommentPosition `json:"origLeftFileEnd,omitempty"`
    // Original position of first character of the thread's span in left file.
    OrigLeftFileStart *CommentPosition `json:"origLeftFileStart,omitempty"`
    // Original position of last character of the thread's span in right file.
    OrigRightFileEnd *CommentPosition `json:"origRightFileEnd,omitempty"`
    // Original position of first character of the thread's span in right file.
    OrigRightFileStart *CommentPosition `json:"origRightFileStart,omitempty"`
    // The iteration of the file on the right side of the diff that the thread will be tracked to. Threads were tracked if this is greater than 0.
    SecondComparingIteration *int `json:"secondComparingIteration,omitempty"`
}

// The type of a comment.
type CommentType string

// Real time event (SignalR) for a completion errors on a pull request
type CompletionErrorsEvent struct {
    // The id of this event. Can be used to track send/receive state between client and server.
    EventId *uuid.UUID `json:"eventId,omitempty"`
    // The id of the pull request this event was generated for.
    PullRequestId *int `json:"pullRequestId,omitempty"`
    // The error message associated with the completion error
    ErrorMessage *string `json:"errorMessage,omitempty"`
}

// Real time event (SignalR) for a discussions update on a pull request
type DiscussionsUpdatedEvent struct {
    // The id of this event. Can be used to track send/receive state between client and server.
    EventId *uuid.UUID `json:"eventId,omitempty"`
    // The id of the pull request this event was generated for.
    PullRequestId *int `json:"pullRequestId,omitempty"`
}

type FileContentMetadata struct {
    ContentType *string `json:"contentType,omitempty"`
    Encoding *int `json:"encoding,omitempty"`
    Extension *string `json:"extension,omitempty"`
    FileName *string `json:"fileName,omitempty"`
    IsBinary *bool `json:"isBinary,omitempty"`
    IsImage *bool `json:"isImage,omitempty"`
    VsLink *string `json:"vsLink,omitempty"`
}

// Provides properties that describe file differences
type FileDiff struct {
    // The collection of line diff blocks
    LineDiffBlocks *[]LineDiffBlock `json:"lineDiffBlocks,omitempty"`
    // Original path of item if different from current path.
    OriginalPath *string `json:"originalPath,omitempty"`
    // Current path of item
    Path *string `json:"path,omitempty"`
}

// Provides parameters that describe inputs for the file diff
type FileDiffParams struct {
    // Original path of the file
    OriginalPath *string `json:"originalPath,omitempty"`
    // Current path of the file
    Path *string `json:"path,omitempty"`
}

// Provides properties that describe inputs for the file diffs
type FileDiffsCriteria struct {
    // Commit ID of the base version
    BaseVersionCommit *string `json:"baseVersionCommit,omitempty"`
    // List of parameters for each of the files for which we need to get the file diff
    FileDiffParams *[]FileDiffParams `json:"fileDiffParams,omitempty"`
    // Commit ID of the target version
    TargetVersionCommit *string `json:"targetVersionCommit,omitempty"`
}

// A Git annotated tag.
type GitAnnotatedTag struct {
    // The tagging Message
    Message *string `json:"message,omitempty"`
    // The name of the annotated tag.
    Name *string `json:"name,omitempty"`
    // The objectId (Sha1Id) of the tag.
    ObjectId *string `json:"objectId,omitempty"`
    // User info and date of tagging.
    TaggedBy *GitUserDate `json:"taggedBy,omitempty"`
    // Tagged git object.
    TaggedObject *GitObject `json:"taggedObject,omitempty"`
    Url *string `json:"url,omitempty"`
}

// Current status of the asynchronous operation.
type GitAsyncOperationStatus string

type GitAsyncRefOperation struct {
    Links *ReferenceLinks `json:"_links,omitempty"`
    DetailedStatus *GitAsyncRefOperationDetail `json:"detailedStatus,omitempty"`
    Parameters *GitAsyncRefOperationParameters `json:"parameters,omitempty"`
    Status *GitAsyncOperationStatus `json:"status,omitempty"`
    // A URL that can be used to make further requests for status about the operation
    Url *string `json:"url,omitempty"`
}

// Information about the progress of a cherry pick or revert operation.
type GitAsyncRefOperationDetail struct {
    // Indicates if there was a conflict generated when trying to cherry pick or revert the changes.
    Conflict *bool `json:"conflict,omitempty"`
    // The current commit from the list of commits that are being cherry picked or reverted.
    CurrentCommitId *string `json:"currentCommitId,omitempty"`
    // Detailed information about why the cherry pick or revert failed to complete.
    FailureMessage *string `json:"failureMessage,omitempty"`
    // A number between 0 and 1 indicating the percent complete of the operation.
    Progress *float64 `json:"progress,omitempty"`
    // Provides a status code that indicates the reason the cherry pick or revert failed.
    Status *GitAsyncRefOperationFailureStatus `json:"status,omitempty"`
    // Indicates if the operation went beyond the maximum time allowed for a cherry pick or revert operation.
    Timedout *bool `json:"timedout,omitempty"`
}

type GitAsyncRefOperationFailureStatus string

// Parameters that are provided in the request body when requesting to cherry pick or revert.
type GitAsyncRefOperationParameters struct {
    // Proposed target branch name for the cherry pick or revert operation.
    GeneratedRefName *string `json:"generatedRefName,omitempty"`
    // The target branch for the cherry pick or revert operation.
    OntoRefName *string `json:"ontoRefName,omitempty"`
    // The git repository for the cherry pick or revert operation.
    Repository *GitRepository `json:"repository,omitempty"`
    // Details about the source of the cherry pick or revert operation (e.g. A pull request or a specific commit).
    Source *GitAsyncRefOperationSource `json:"source,omitempty"`
}

// GitAsyncRefOperationSource specifies the pull request or list of commits to use when making a cherry pick and revert operation request. Only one should be provided.
type GitAsyncRefOperationSource struct {
    // A list of commits to cherry pick or revert
    CommitList *[]GitCommitRef `json:"commitList,omitempty"`
    // Id of the pull request to cherry pick or revert
    PullRequestId *int `json:"pullRequestId,omitempty"`
}

type GitBaseVersionDescriptor struct {
    // Version string identifier (name of tag/branch, SHA1 of commit)
    Version *string `json:"version,omitempty"`
    // Version options - Specify additional modifiers to version (e.g Previous)
    VersionOptions *GitVersionOptions `json:"versionOptions,omitempty"`
    // Version type (branch, tag, or commit). Determines how Id is interpreted
    VersionType *GitVersionType `json:"versionType,omitempty"`
    // Version string identifier (name of tag/branch, SHA1 of commit)
    BaseVersion *string `json:"baseVersion,omitempty"`
    // Version options - Specify additional modifiers to version (e.g Previous)
    BaseVersionOptions *GitVersionOptions `json:"baseVersionOptions,omitempty"`
    // Version type (branch, tag, or commit). Determines how Id is interpreted
    BaseVersionType *GitVersionType `json:"baseVersionType,omitempty"`
}

type GitBlobRef struct {
    Links *ReferenceLinks `json:"_links,omitempty"`
    // SHA1 hash of git object
    ObjectId *string `json:"objectId,omitempty"`
    // Size of blob content (in bytes)
    Size *uint64 `json:"size,omitempty"`
    Url *string `json:"url,omitempty"`
}

// Ahead and behind counts for a particular ref.
type GitBranchStats struct {
    // Number of commits ahead.
    AheadCount *int `json:"aheadCount,omitempty"`
    // Number of commits behind.
    BehindCount *int `json:"behindCount,omitempty"`
    // Current commit.
    Commit *GitCommitRef `json:"commit,omitempty"`
    // True if this is the result for the base version.
    IsBaseVersion *bool `json:"isBaseVersion,omitempty"`
    // Name of the ref.
    Name *string `json:"name,omitempty"`
}

// This object is returned from Cherry Pick operations and provides the id and status of the operation
type GitCherryPick struct {
    Links *ReferenceLinks `json:"_links,omitempty"`
    DetailedStatus *GitAsyncRefOperationDetail `json:"detailedStatus,omitempty"`
    Parameters *GitAsyncRefOperationParameters `json:"parameters,omitempty"`
    Status *GitAsyncOperationStatus `json:"status,omitempty"`
    // A URL that can be used to make further requests for status about the operation
    Url *string `json:"url,omitempty"`
    CherryPickId *int `json:"cherryPickId,omitempty"`
}

type GitCommit struct {
    // A collection of related REST reference links.
    Links *ReferenceLinks `json:"_links,omitempty"`
    // Author of the commit.
    Author *GitUserDate `json:"author,omitempty"`
    // Counts of the types of changes (edits, deletes, etc.) included with the commit.
    ChangeCounts *ChangeCountDictionary `json:"changeCounts,omitempty"`
    // An enumeration of the changes included with the commit.
    Changes *[]interface{} `json:"changes,omitempty"`
    // Comment or message of the commit.
    Comment *string `json:"comment,omitempty"`
    // Indicates if the comment is truncated from the full Git commit comment message.
    CommentTruncated *bool `json:"commentTruncated,omitempty"`
    // ID (SHA-1) of the commit.
    CommitId *string `json:"commitId,omitempty"`
    // Committer of the commit.
    Committer *GitUserDate `json:"committer,omitempty"`
    // An enumeration of the parent commit IDs for this commit.
    Parents *[]string `json:"parents,omitempty"`
    // The push associated with this commit.
    Push *GitPushRef `json:"push,omitempty"`
    // Remote URL path to the commit.
    RemoteUrl *string `json:"remoteUrl,omitempty"`
    // A list of status metadata from services and extensions that may associate additional information to the commit.
    Statuses *[]GitStatus `json:"statuses,omitempty"`
    // REST URL for this resource.
    Url *string `json:"url,omitempty"`
    // A list of workitems associated with this commit.
    WorkItems *[]ResourceRef `json:"workItems,omitempty"`
    TreeId *string `json:"treeId,omitempty"`
}

type GitCommitChanges struct {
    ChangeCounts *ChangeCountDictionary `json:"changeCounts,omitempty"`
    Changes *[]interface{} `json:"changes,omitempty"`
}

type GitCommitDiffs struct {
    AheadCount *int `json:"aheadCount,omitempty"`
    AllChangesIncluded *bool `json:"allChangesIncluded,omitempty"`
    BaseCommit *string `json:"baseCommit,omitempty"`
    BehindCount *int `json:"behindCount,omitempty"`
    ChangeCounts *map[VersionControlChangeType]int `json:"changeCounts,omitempty"`
    Changes *[]interface{} `json:"changes,omitempty"`
    CommonCommit *string `json:"commonCommit,omitempty"`
    TargetCommit *string `json:"targetCommit,omitempty"`
}

// Provides properties that describe a Git commit and associated metadata.
type GitCommitRef struct {
    // A collection of related REST reference links.
    Links *ReferenceLinks `json:"_links,omitempty"`
    // Author of the commit.
    Author *GitUserDate `json:"author,omitempty"`
    // Counts of the types of changes (edits, deletes, etc.) included with the commit.
    ChangeCounts *ChangeCountDictionary `json:"changeCounts,omitempty"`
    // An enumeration of the changes included with the commit.
    Changes *[]interface{} `json:"changes,omitempty"`
    // Comment or message of the commit.
    Comment *string `json:"comment,omitempty"`
    // Indicates if the comment is truncated from the full Git commit comment message.
    CommentTruncated *bool `json:"commentTruncated,omitempty"`
    // ID (SHA-1) of the commit.
    CommitId *string `json:"commitId,omitempty"`
    // Committer of the commit.
    Committer *GitUserDate `json:"committer,omitempty"`
    // An enumeration of the parent commit IDs for this commit.
    Parents *[]string `json:"parents,omitempty"`
    // The push associated with this commit.
    Push *GitPushRef `json:"push,omitempty"`
    // Remote URL path to the commit.
    RemoteUrl *string `json:"remoteUrl,omitempty"`
    // A list of status metadata from services and extensions that may associate additional information to the commit.
    Statuses *[]GitStatus `json:"statuses,omitempty"`
    // REST URL for this resource.
    Url *string `json:"url,omitempty"`
    // A list of workitems associated with this commit.
    WorkItems *[]ResourceRef `json:"workItems,omitempty"`
}

type GitCommitToCreate struct {
    BaseRef *GitRef `json:"baseRef,omitempty"`
    Comment *string `json:"comment,omitempty"`
    PathActions *[]GitPathAction `json:"pathActions,omitempty"`
}

type GitConflict struct {
    Links *ReferenceLinks `json:"_links,omitempty"`
    ConflictId *int `json:"conflictId,omitempty"`
    ConflictPath *string `json:"conflictPath,omitempty"`
    ConflictType *GitConflictType `json:"conflictType,omitempty"`
    MergeBaseCommit *GitCommitRef `json:"mergeBaseCommit,omitempty"`
    MergeOrigin *GitMergeOriginRef `json:"mergeOrigin,omitempty"`
    MergeSourceCommit *GitCommitRef `json:"mergeSourceCommit,omitempty"`
    MergeTargetCommit *GitCommitRef `json:"mergeTargetCommit,omitempty"`
    ResolutionError *GitResolutionError `json:"resolutionError,omitempty"`
    ResolutionStatus *GitResolutionStatus `json:"resolutionStatus,omitempty"`
    ResolvedBy *IdentityRef `json:"resolvedBy,omitempty"`
    ResolvedDate *time.Time `json:"resolvedDate,omitempty"`
    Url *string `json:"url,omitempty"`
}

// Data object for AddAdd conflict
type GitConflictAddAdd struct {
    Links *ReferenceLinks `json:"_links,omitempty"`
    ConflictId *int `json:"conflictId,omitempty"`
    ConflictPath *string `json:"conflictPath,omitempty"`
    ConflictType *GitConflictType `json:"conflictType,omitempty"`
    MergeBaseCommit *GitCommitRef `json:"mergeBaseCommit,omitempty"`
    MergeOrigin *GitMergeOriginRef `json:"mergeOrigin,omitempty"`
    MergeSourceCommit *GitCommitRef `json:"mergeSourceCommit,omitempty"`
    MergeTargetCommit *GitCommitRef `json:"mergeTargetCommit,omitempty"`
    ResolutionError *GitResolutionError `json:"resolutionError,omitempty"`
    ResolutionStatus *GitResolutionStatus `json:"resolutionStatus,omitempty"`
    ResolvedBy *IdentityRef `json:"resolvedBy,omitempty"`
    ResolvedDate *time.Time `json:"resolvedDate,omitempty"`
    Url *string `json:"url,omitempty"`
    Resolution *GitResolutionMergeContent `json:"resolution,omitempty"`
    SourceBlob *GitBlobRef `json:"sourceBlob,omitempty"`
    TargetBlob *GitBlobRef `json:"targetBlob,omitempty"`
}

// Data object for RenameAdd conflict
type GitConflictAddRename struct {
    Links *ReferenceLinks `json:"_links,omitempty"`
    ConflictId *int `json:"conflictId,omitempty"`
    ConflictPath *string `json:"conflictPath,omitempty"`
    ConflictType *GitConflictType `json:"conflictType,omitempty"`
    MergeBaseCommit *GitCommitRef `json:"mergeBaseCommit,omitempty"`
    MergeOrigin *GitMergeOriginRef `json:"mergeOrigin,omitempty"`
    MergeSourceCommit *GitCommitRef `json:"mergeSourceCommit,omitempty"`
    MergeTargetCommit *GitCommitRef `json:"mergeTargetCommit,omitempty"`
    ResolutionError *GitResolutionError `json:"resolutionError,omitempty"`
    ResolutionStatus *GitResolutionStatus `json:"resolutionStatus,omitempty"`
    ResolvedBy *IdentityRef `json:"resolvedBy,omitempty"`
    ResolvedDate *time.Time `json:"resolvedDate,omitempty"`
    Url *string `json:"url,omitempty"`
    BaseBlob *GitBlobRef `json:"baseBlob,omitempty"`
    Resolution *GitResolutionPathConflict `json:"resolution,omitempty"`
    SourceBlob *GitBlobRef `json:"sourceBlob,omitempty"`
    TargetBlob *GitBlobRef `json:"targetBlob,omitempty"`
    TargetOriginalPath *string `json:"targetOriginalPath,omitempty"`
}

// Data object for EditDelete conflict
type GitConflictDeleteEdit struct {
    Links *ReferenceLinks `json:"_links,omitempty"`
    ConflictId *int `json:"conflictId,omitempty"`
    ConflictPath *string `json:"conflictPath,omitempty"`
    ConflictType *GitConflictType `json:"conflictType,omitempty"`
    MergeBaseCommit *GitCommitRef `json:"mergeBaseCommit,omitempty"`
    MergeOrigin *GitMergeOriginRef `json:"mergeOrigin,omitempty"`
    MergeSourceCommit *GitCommitRef `json:"mergeSourceCommit,omitempty"`
    MergeTargetCommit *GitCommitRef `json:"mergeTargetCommit,omitempty"`
    ResolutionError *GitResolutionError `json:"resolutionError,omitempty"`
    ResolutionStatus *GitResolutionStatus `json:"resolutionStatus,omitempty"`
    ResolvedBy *IdentityRef `json:"resolvedBy,omitempty"`
    ResolvedDate *time.Time `json:"resolvedDate,omitempty"`
    Url *string `json:"url,omitempty"`
    BaseBlob *GitBlobRef `json:"baseBlob,omitempty"`
    Resolution *GitResolutionPickOneAction `json:"resolution,omitempty"`
    TargetBlob *GitBlobRef `json:"targetBlob,omitempty"`
}

// Data object for RenameDelete conflict
type GitConflictDeleteRename struct {
    Links *ReferenceLinks `json:"_links,omitempty"`
    ConflictId *int `json:"conflictId,omitempty"`
    ConflictPath *string `json:"conflictPath,omitempty"`
    ConflictType *GitConflictType `json:"conflictType,omitempty"`
    MergeBaseCommit *GitCommitRef `json:"mergeBaseCommit,omitempty"`
    MergeOrigin *GitMergeOriginRef `json:"mergeOrigin,omitempty"`
    MergeSourceCommit *GitCommitRef `json:"mergeSourceCommit,omitempty"`
    MergeTargetCommit *GitCommitRef `json:"mergeTargetCommit,omitempty"`
    ResolutionError *GitResolutionError `json:"resolutionError,omitempty"`
    ResolutionStatus *GitResolutionStatus `json:"resolutionStatus,omitempty"`
    ResolvedBy *IdentityRef `json:"resolvedBy,omitempty"`
    ResolvedDate *time.Time `json:"resolvedDate,omitempty"`
    Url *string `json:"url,omitempty"`
    BaseBlob *GitBlobRef `json:"baseBlob,omitempty"`
    Resolution *GitResolutionPickOneAction `json:"resolution,omitempty"`
    TargetBlob *GitBlobRef `json:"targetBlob,omitempty"`
    TargetNewPath *string `json:"targetNewPath,omitempty"`
}

// Data object for FileDirectory conflict
type GitConflictDirectoryFile struct {
    Links *ReferenceLinks `json:"_links,omitempty"`
    ConflictId *int `json:"conflictId,omitempty"`
    ConflictPath *string `json:"conflictPath,omitempty"`
    ConflictType *GitConflictType `json:"conflictType,omitempty"`
    MergeBaseCommit *GitCommitRef `json:"mergeBaseCommit,omitempty"`
    MergeOrigin *GitMergeOriginRef `json:"mergeOrigin,omitempty"`
    MergeSourceCommit *GitCommitRef `json:"mergeSourceCommit,omitempty"`
    MergeTargetCommit *GitCommitRef `json:"mergeTargetCommit,omitempty"`
    ResolutionError *GitResolutionError `json:"resolutionError,omitempty"`
    ResolutionStatus *GitResolutionStatus `json:"resolutionStatus,omitempty"`
    ResolvedBy *IdentityRef `json:"resolvedBy,omitempty"`
    ResolvedDate *time.Time `json:"resolvedDate,omitempty"`
    Url *string `json:"url,omitempty"`
    Resolution *GitResolutionPathConflict `json:"resolution,omitempty"`
    SourceTree *GitTreeRef `json:"sourceTree,omitempty"`
    TargetBlob *GitBlobRef `json:"targetBlob,omitempty"`
}

// Data object for DeleteEdit conflict
type GitConflictEditDelete struct {
    Links *ReferenceLinks `json:"_links,omitempty"`
    ConflictId *int `json:"conflictId,omitempty"`
    ConflictPath *string `json:"conflictPath,omitempty"`
    ConflictType *GitConflictType `json:"conflictType,omitempty"`
    MergeBaseCommit *GitCommitRef `json:"mergeBaseCommit,omitempty"`
    MergeOrigin *GitMergeOriginRef `json:"mergeOrigin,omitempty"`
    MergeSourceCommit *GitCommitRef `json:"mergeSourceCommit,omitempty"`
    MergeTargetCommit *GitCommitRef `json:"mergeTargetCommit,omitempty"`
    ResolutionError *GitResolutionError `json:"resolutionError,omitempty"`
    ResolutionStatus *GitResolutionStatus `json:"resolutionStatus,omitempty"`
    ResolvedBy *IdentityRef `json:"resolvedBy,omitempty"`
    ResolvedDate *time.Time `json:"resolvedDate,omitempty"`
    Url *string `json:"url,omitempty"`
    BaseBlob *GitBlobRef `json:"baseBlob,omitempty"`
    Resolution *GitResolutionPickOneAction `json:"resolution,omitempty"`
    SourceBlob *GitBlobRef `json:"sourceBlob,omitempty"`
}

// Data object for EditEdit conflict
type GitConflictEditEdit struct {
    Links *ReferenceLinks `json:"_links,omitempty"`
    ConflictId *int `json:"conflictId,omitempty"`
    ConflictPath *string `json:"conflictPath,omitempty"`
    ConflictType *GitConflictType `json:"conflictType,omitempty"`
    MergeBaseCommit *GitCommitRef `json:"mergeBaseCommit,omitempty"`
    MergeOrigin *GitMergeOriginRef `json:"mergeOrigin,omitempty"`
    MergeSourceCommit *GitCommitRef `json:"mergeSourceCommit,omitempty"`
    MergeTargetCommit *GitCommitRef `json:"mergeTargetCommit,omitempty"`
    ResolutionError *GitResolutionError `json:"resolutionError,omitempty"`
    ResolutionStatus *GitResolutionStatus `json:"resolutionStatus,omitempty"`
    ResolvedBy *IdentityRef `json:"resolvedBy,omitempty"`
    ResolvedDate *time.Time `json:"resolvedDate,omitempty"`
    Url *string `json:"url,omitempty"`
    BaseBlob *GitBlobRef `json:"baseBlob,omitempty"`
    Resolution *GitResolutionMergeContent `json:"resolution,omitempty"`
    SourceBlob *GitBlobRef `json:"sourceBlob,omitempty"`
    TargetBlob *GitBlobRef `json:"targetBlob,omitempty"`
}

// Data object for DirectoryFile conflict
type GitConflictFileDirectory struct {
    Links *ReferenceLinks `json:"_links,omitempty"`
    ConflictId *int `json:"conflictId,omitempty"`
    ConflictPath *string `json:"conflictPath,omitempty"`
    ConflictType *GitConflictType `json:"conflictType,omitempty"`
    MergeBaseCommit *GitCommitRef `json:"mergeBaseCommit,omitempty"`
    MergeOrigin *GitMergeOriginRef `json:"mergeOrigin,omitempty"`
    MergeSourceCommit *GitCommitRef `json:"mergeSourceCommit,omitempty"`
    MergeTargetCommit *GitCommitRef `json:"mergeTargetCommit,omitempty"`
    ResolutionError *GitResolutionError `json:"resolutionError,omitempty"`
    ResolutionStatus *GitResolutionStatus `json:"resolutionStatus,omitempty"`
    ResolvedBy *IdentityRef `json:"resolvedBy,omitempty"`
    ResolvedDate *time.Time `json:"resolvedDate,omitempty"`
    Url *string `json:"url,omitempty"`
    Resolution *GitResolutionPathConflict `json:"resolution,omitempty"`
    SourceBlob *GitBlobRef `json:"sourceBlob,omitempty"`
    TargetTree *GitTreeRef `json:"targetTree,omitempty"`
}

// Data object for Rename1to2 conflict
type GitConflictRename1to2 struct {
    Links *ReferenceLinks `json:"_links,omitempty"`
    ConflictId *int `json:"conflictId,omitempty"`
    ConflictPath *string `json:"conflictPath,omitempty"`
    ConflictType *GitConflictType `json:"conflictType,omitempty"`
    MergeBaseCommit *GitCommitRef `json:"mergeBaseCommit,omitempty"`
    MergeOrigin *GitMergeOriginRef `json:"mergeOrigin,omitempty"`
    MergeSourceCommit *GitCommitRef `json:"mergeSourceCommit,omitempty"`
    MergeTargetCommit *GitCommitRef `json:"mergeTargetCommit,omitempty"`
    ResolutionError *GitResolutionError `json:"resolutionError,omitempty"`
    ResolutionStatus *GitResolutionStatus `json:"resolutionStatus,omitempty"`
    ResolvedBy *IdentityRef `json:"resolvedBy,omitempty"`
    ResolvedDate *time.Time `json:"resolvedDate,omitempty"`
    Url *string `json:"url,omitempty"`
    BaseBlob *GitBlobRef `json:"baseBlob,omitempty"`
    Resolution *GitResolutionRename1to2 `json:"resolution,omitempty"`
    SourceBlob *GitBlobRef `json:"sourceBlob,omitempty"`
    SourceNewPath *string `json:"sourceNewPath,omitempty"`
    TargetBlob *GitBlobRef `json:"targetBlob,omitempty"`
    TargetNewPath *string `json:"targetNewPath,omitempty"`
}

// Data object for Rename2to1 conflict
type GitConflictRename2to1 struct {
    Links *ReferenceLinks `json:"_links,omitempty"`
    ConflictId *int `json:"conflictId,omitempty"`
    ConflictPath *string `json:"conflictPath,omitempty"`
    ConflictType *GitConflictType `json:"conflictType,omitempty"`
    MergeBaseCommit *GitCommitRef `json:"mergeBaseCommit,omitempty"`
    MergeOrigin *GitMergeOriginRef `json:"mergeOrigin,omitempty"`
    MergeSourceCommit *GitCommitRef `json:"mergeSourceCommit,omitempty"`
    MergeTargetCommit *GitCommitRef `json:"mergeTargetCommit,omitempty"`
    ResolutionError *GitResolutionError `json:"resolutionError,omitempty"`
    ResolutionStatus *GitResolutionStatus `json:"resolutionStatus,omitempty"`
    ResolvedBy *IdentityRef `json:"resolvedBy,omitempty"`
    ResolvedDate *time.Time `json:"resolvedDate,omitempty"`
    Url *string `json:"url,omitempty"`
    Resolution *GitResolutionPathConflict `json:"resolution,omitempty"`
    SourceNewBlob *GitBlobRef `json:"sourceNewBlob,omitempty"`
    SourceOriginalBlob *GitBlobRef `json:"sourceOriginalBlob,omitempty"`
    SourceOriginalPath *string `json:"sourceOriginalPath,omitempty"`
    TargetNewBlob *GitBlobRef `json:"targetNewBlob,omitempty"`
    TargetOriginalBlob *GitBlobRef `json:"targetOriginalBlob,omitempty"`
    TargetOriginalPath *string `json:"targetOriginalPath,omitempty"`
}

// Data object for AddRename conflict
type GitConflictRenameAdd struct {
    Links *ReferenceLinks `json:"_links,omitempty"`
    ConflictId *int `json:"conflictId,omitempty"`
    ConflictPath *string `json:"conflictPath,omitempty"`
    ConflictType *GitConflictType `json:"conflictType,omitempty"`
    MergeBaseCommit *GitCommitRef `json:"mergeBaseCommit,omitempty"`
    MergeOrigin *GitMergeOriginRef `json:"mergeOrigin,omitempty"`
    MergeSourceCommit *GitCommitRef `json:"mergeSourceCommit,omitempty"`
    MergeTargetCommit *GitCommitRef `json:"mergeTargetCommit,omitempty"`
    ResolutionError *GitResolutionError `json:"resolutionError,omitempty"`
    ResolutionStatus *GitResolutionStatus `json:"resolutionStatus,omitempty"`
    ResolvedBy *IdentityRef `json:"resolvedBy,omitempty"`
    ResolvedDate *time.Time `json:"resolvedDate,omitempty"`
    Url *string `json:"url,omitempty"`
    BaseBlob *GitBlobRef `json:"baseBlob,omitempty"`
    Resolution *GitResolutionPathConflict `json:"resolution,omitempty"`
    SourceBlob *GitBlobRef `json:"sourceBlob,omitempty"`
    SourceOriginalPath *string `json:"sourceOriginalPath,omitempty"`
    TargetBlob *GitBlobRef `json:"targetBlob,omitempty"`
}

// Data object for DeleteRename conflict
type GitConflictRenameDelete struct {
    Links *ReferenceLinks `json:"_links,omitempty"`
    ConflictId *int `json:"conflictId,omitempty"`
    ConflictPath *string `json:"conflictPath,omitempty"`
    ConflictType *GitConflictType `json:"conflictType,omitempty"`
    MergeBaseCommit *GitCommitRef `json:"mergeBaseCommit,omitempty"`
    MergeOrigin *GitMergeOriginRef `json:"mergeOrigin,omitempty"`
    MergeSourceCommit *GitCommitRef `json:"mergeSourceCommit,omitempty"`
    MergeTargetCommit *GitCommitRef `json:"mergeTargetCommit,omitempty"`
    ResolutionError *GitResolutionError `json:"resolutionError,omitempty"`
    ResolutionStatus *GitResolutionStatus `json:"resolutionStatus,omitempty"`
    ResolvedBy *IdentityRef `json:"resolvedBy,omitempty"`
    ResolvedDate *time.Time `json:"resolvedDate,omitempty"`
    Url *string `json:"url,omitempty"`
    BaseBlob *GitBlobRef `json:"baseBlob,omitempty"`
    Resolution *GitResolutionPickOneAction `json:"resolution,omitempty"`
    SourceBlob *GitBlobRef `json:"sourceBlob,omitempty"`
    SourceNewPath *string `json:"sourceNewPath,omitempty"`
}

// Data object for RenameRename conflict
type GitConflictRenameRename struct {
    Links *ReferenceLinks `json:"_links,omitempty"`
    ConflictId *int `json:"conflictId,omitempty"`
    ConflictPath *string `json:"conflictPath,omitempty"`
    ConflictType *GitConflictType `json:"conflictType,omitempty"`
    MergeBaseCommit *GitCommitRef `json:"mergeBaseCommit,omitempty"`
    MergeOrigin *GitMergeOriginRef `json:"mergeOrigin,omitempty"`
    MergeSourceCommit *GitCommitRef `json:"mergeSourceCommit,omitempty"`
    MergeTargetCommit *GitCommitRef `json:"mergeTargetCommit,omitempty"`
    ResolutionError *GitResolutionError `json:"resolutionError,omitempty"`
    ResolutionStatus *GitResolutionStatus `json:"resolutionStatus,omitempty"`
    ResolvedBy *IdentityRef `json:"resolvedBy,omitempty"`
    ResolvedDate *time.Time `json:"resolvedDate,omitempty"`
    Url *string `json:"url,omitempty"`
    BaseBlob *GitBlobRef `json:"baseBlob,omitempty"`
    OriginalPath *string `json:"originalPath,omitempty"`
    Resolution *GitResolutionMergeContent `json:"resolution,omitempty"`
    SourceBlob *GitBlobRef `json:"sourceBlob,omitempty"`
    TargetBlob *GitBlobRef `json:"targetBlob,omitempty"`
}

// The type of a merge conflict.
type GitConflictType string

type GitConflictUpdateResult struct {
    // Conflict ID that was provided by input
    ConflictId *int `json:"conflictId,omitempty"`
    // Reason for failing
    CustomMessage *string `json:"customMessage,omitempty"`
    // New state of the conflict after updating
    UpdatedConflict *GitConflict `json:"updatedConflict,omitempty"`
    // Status of the update on the server
    UpdateStatus *GitConflictUpdateStatus `json:"updateStatus,omitempty"`
}

// Represents the possible outcomes from a request to update a pull request conflict
type GitConflictUpdateStatus string

type GitDeletedRepository struct {
    CreatedDate *time.Time `json:"createdDate,omitempty"`
    DeletedBy *IdentityRef `json:"deletedBy,omitempty"`
    DeletedDate *time.Time `json:"deletedDate,omitempty"`
    Id *uuid.UUID `json:"id,omitempty"`
    Name *string `json:"name,omitempty"`
    Project *TeamProjectReference `json:"project,omitempty"`
}

type GitFilePathsCollection struct {
    CommitId *string `json:"commitId,omitempty"`
    Paths *[]string `json:"paths,omitempty"`
    Url *string `json:"url,omitempty"`
}

// Status information about a requested fork operation.
type GitForkOperationStatusDetail struct {
    // All valid steps for the forking process
    AllSteps *[]string `json:"allSteps,omitempty"`
    // Index into AllSteps for the current step
    CurrentStep *int `json:"currentStep,omitempty"`
    // Error message if the operation failed.
    ErrorMessage *string `json:"errorMessage,omitempty"`
}

// Information about a fork ref.
type GitForkRef struct {
    Links *ReferenceLinks `json:"_links,omitempty"`
    Creator *IdentityRef `json:"creator,omitempty"`
    IsLocked *bool `json:"isLocked,omitempty"`
    IsLockedBy *IdentityRef `json:"isLockedBy,omitempty"`
    Name *string `json:"name,omitempty"`
    ObjectId *string `json:"objectId,omitempty"`
    PeeledObjectId *string `json:"peeledObjectId,omitempty"`
    Statuses *[]GitStatus `json:"statuses,omitempty"`
    Url *string `json:"url,omitempty"`
    // The repository ID of the fork.
    Repository *GitRepository `json:"repository,omitempty"`
}

// Request to sync data between two forks.
type GitForkSyncRequest struct {
    // Collection of related links
    Links *ReferenceLinks `json:"_links,omitempty"`
    DetailedStatus *GitForkOperationStatusDetail `json:"detailedStatus,omitempty"`
    // Unique identifier for the operation.
    OperationId *int `json:"operationId,omitempty"`
    // Fully-qualified identifier for the source repository.
    Source *GlobalGitRepositoryKey `json:"source,omitempty"`
    // If supplied, the set of ref mappings to use when performing a "sync" or create. If missing, all refs will be synchronized.
    SourceToTargetRefs *[]SourceToTargetRef `json:"sourceToTargetRefs,omitempty"`
    Status *GitAsyncOperationStatus `json:"status,omitempty"`
}

// Parameters for creating a fork request
type GitForkSyncRequestParameters struct {
    // Fully-qualified identifier for the source repository.
    Source *GlobalGitRepositoryKey `json:"source,omitempty"`
    // If supplied, the set of ref mappings to use when performing a "sync" or create. If missing, all refs will be synchronized.
    SourceToTargetRefs *[]SourceToTargetRef `json:"sourceToTargetRefs,omitempty"`
}

type GitForkTeamProjectReference struct {
    // Project abbreviation.
    Abbreviation *string `json:"abbreviation,omitempty"`
    // Url to default team identity image.
    DefaultTeamImageUrl *string `json:"defaultTeamImageUrl,omitempty"`
    // The project's description (if any).
    Description *string `json:"description,omitempty"`
    // Project identifier.
    Id *uuid.UUID `json:"id,omitempty"`
    // Project last update time.
    LastUpdateTime *time.Time `json:"lastUpdateTime,omitempty"`
    // Project name.
    Name *string `json:"name,omitempty"`
    // Project revision.
    Revision *uint64 `json:"revision,omitempty"`
    // Project state.
    State *interface{} `json:"state,omitempty"`
    // Url to the full version of the object.
    Url *string `json:"url,omitempty"`
    // Project visibility.
    Visibility *ProjectVisibility `json:"visibility,omitempty"`
}

// Accepted types of version
type GitHistoryMode string

type GitImportFailedEvent struct {
    SourceRepositoryName *string `json:"sourceRepositoryName,omitempty"`
    TargetRepository *GitRepository `json:"targetRepository,omitempty"`
}

// Parameter for creating a git import request when source is Git version control
type GitImportGitSource struct {
    // Tells if this is a sync request or not
    Overwrite *bool `json:"overwrite,omitempty"`
    // Url for the source repo
    Url *string `json:"url,omitempty"`
}

// A request to import data from a remote source control system.
type GitImportRequest struct {
    // Links to related resources.
    Links *ReferenceLinks `json:"_links,omitempty"`
    // Detailed status of the import, including the current step and an error message, if applicable.
    DetailedStatus *GitImportStatusDetail `json:"detailedStatus,omitempty"`
    // The unique identifier for this import request.
    ImportRequestId *int `json:"importRequestId,omitempty"`
    // Parameters for creating the import request.
    Parameters *GitImportRequestParameters `json:"parameters,omitempty"`
    // The target repository for this import.
    Repository *GitRepository `json:"repository,omitempty"`
    // Current status of the import.
    Status *GitAsyncOperationStatus `json:"status,omitempty"`
    // A link back to this import request resource.
    Url *string `json:"url,omitempty"`
}

// Parameters for creating an import request
type GitImportRequestParameters struct {
    // Option to delete service endpoint when import is done
    DeleteServiceEndpointAfterImportIsDone *bool `json:"deleteServiceEndpointAfterImportIsDone,omitempty"`
    // Source for importing git repository
    GitSource *GitImportGitSource `json:"gitSource,omitempty"`
    // Service Endpoint for connection to external endpoint
    ServiceEndpointId *uuid.UUID `json:"serviceEndpointId,omitempty"`
    // Source for importing tfvc repository
    TfvcSource *GitImportTfvcSource `json:"tfvcSource,omitempty"`
}

// Additional status information about an import request.
type GitImportStatusDetail struct {
    // All valid steps for the import process
    AllSteps *[]string `json:"allSteps,omitempty"`
    // Index into AllSteps for the current step
    CurrentStep *int `json:"currentStep,omitempty"`
    // Error message if the operation failed.
    ErrorMessage *string `json:"errorMessage,omitempty"`
}

type GitImportSucceededEvent struct {
    SourceRepositoryName *string `json:"sourceRepositoryName,omitempty"`
    TargetRepository *GitRepository `json:"targetRepository,omitempty"`
}

// Parameter for creating a git import request when source is tfvc version control
type GitImportTfvcSource struct {
    // Set true to import History, false otherwise
    ImportHistory *bool `json:"importHistory,omitempty"`
    // Get history for last n days (max allowed value is 180 days)
    ImportHistoryDurationInDays *int `json:"importHistoryDurationInDays,omitempty"`
    // Path which we want to import (this can be copied from Path Control in Explorer)
    Path *string `json:"path,omitempty"`
}

type GitItem struct {
    Links *ReferenceLinks `json:"_links,omitempty"`
    Content *string `json:"content,omitempty"`
    ContentMetadata *FileContentMetadata `json:"contentMetadata,omitempty"`
    IsFolder *bool `json:"isFolder,omitempty"`
    IsSymLink *bool `json:"isSymLink,omitempty"`
    Path *string `json:"path,omitempty"`
    Url *string `json:"url,omitempty"`
    // SHA1 of commit item was fetched at
    CommitId *string `json:"commitId,omitempty"`
    // Type of object (Commit, Tree, Blob, Tag, ...)
    GitObjectType *GitObjectType `json:"gitObjectType,omitempty"`
    // Shallow ref to commit that last changed this item Only populated if latestProcessedChange is requested May not be accurate if latest change is not yet cached
    LatestProcessedChange *GitCommitRef `json:"latestProcessedChange,omitempty"`
    // Git object id
    ObjectId *string `json:"objectId,omitempty"`
    // Git object id
    OriginalObjectId *string `json:"originalObjectId,omitempty"`
}

type GitItemDescriptor struct {
    // Path to item
    Path *string `json:"path,omitempty"`
    // Specifies whether to include children (OneLevel), all descendants (Full), or None
    RecursionLevel *VersionControlRecursionType `json:"recursionLevel,omitempty"`
    // Version string (interpretation based on VersionType defined in subclass
    Version *string `json:"version,omitempty"`
    // Version modifiers (e.g. previous)
    VersionOptions *GitVersionOptions `json:"versionOptions,omitempty"`
    // How to interpret version (branch,tag,commit)
    VersionType *GitVersionType `json:"versionType,omitempty"`
}

type GitItemRequestData struct {
    // Whether to include metadata for all items
    IncludeContentMetadata *bool `json:"includeContentMetadata,omitempty"`
    // Whether to include the _links field on the shallow references
    IncludeLinks *bool `json:"includeLinks,omitempty"`
    // Collection of items to fetch, including path, version, and recursion level
    ItemDescriptors *[]GitItemDescriptor `json:"itemDescriptors,omitempty"`
    // Whether to include shallow ref to commit that last changed each item
    LatestProcessedChange *bool `json:"latestProcessedChange,omitempty"`
}

type GitLastChangeItem struct {
    // Gets or sets the commit Id this item was modified most recently for the provided version.
    CommitId *string `json:"commitId,omitempty"`
    // Gets or sets the path of the item.
    Path *string `json:"path,omitempty"`
}

type GitLastChangeTreeItems struct {
    // The list of commits referenced by Items, if they were requested.
    Commits *[]GitCommitRef `json:"commits,omitempty"`
    // The last change of items.
    Items *[]GitLastChangeItem `json:"items,omitempty"`
    // The last explored time, in case the result is not comprehensive. Null otherwise.
    LastExploredTime *time.Time `json:"lastExploredTime,omitempty"`
}

type GitMerge struct {
    // Comment or message of the commit.
    Comment *string `json:"comment,omitempty"`
    // An enumeration of the parent commit IDs for the merge  commit.
    Parents *[]string `json:"parents,omitempty"`
    // Reference links.
    Links *ReferenceLinks `json:"_links,omitempty"`
    // Detailed status of the merge operation.
    DetailedStatus *GitMergeOperationStatusDetail `json:"detailedStatus,omitempty"`
    // Unique identifier for the merge operation.
    MergeOperationId *int `json:"mergeOperationId,omitempty"`
    // Status of the merge operation.
    Status *GitAsyncOperationStatus `json:"status,omitempty"`
}

// Status information about a requested merge operation.
type GitMergeOperationStatusDetail struct {
    // Error message if the operation failed.
    FailureMessage *string `json:"failureMessage,omitempty"`
    // The commitId of the resultant merge commit.
    MergeCommitId *string `json:"mergeCommitId,omitempty"`
}

type GitMergeOriginRef struct {
    PullRequestId *int `json:"pullRequestId,omitempty"`
}

// Parameters required for performing git merge.
type GitMergeParameters struct {
    // Comment or message of the commit.
    Comment *string `json:"comment,omitempty"`
    // An enumeration of the parent commit IDs for the merge  commit.
    Parents *[]string `json:"parents,omitempty"`
}

// Git object identifier and type information.
type GitObject struct {
    // Object Id (Sha1Id).
    ObjectId *string `json:"objectId,omitempty"`
    // Type of object (Commit, Tree, Blob, Tag)
    ObjectType *GitObjectType `json:"objectType,omitempty"`
}

type GitObjectType string

type GitPathAction struct {
    Action *GitPathActions `json:"action,omitempty"`
    Base64Content *string `json:"base64Content,omitempty"`
    Path *string `json:"path,omitempty"`
    RawTextContent *string `json:"rawTextContent,omitempty"`
    TargetPath *string `json:"targetPath,omitempty"`
}

type GitPathActions string

type GitPathToItemsCollection struct {
    Items *map[string][]GitItem `json:"items,omitempty"`
}

type GitPolicyConfigurationResponse struct {
    // The HTTP client methods find the continuation token header in the response and populate this field.
    ContinuationToken *string `json:"continuationToken,omitempty"`
    PolicyConfigurations *[]PolicyConfiguration `json:"policyConfigurations,omitempty"`
}

// Represents all the data associated with a pull request.
type GitPullRequest struct {
    // Links to other related objects.
    Links *ReferenceLinks `json:"_links,omitempty"`
    // A string which uniquely identifies this pull request. To generate an artifact ID for a pull request, use this template: ```vstfs:///Git/PullRequestId/{projectId}/{repositoryId}/{pullRequestId}```
    ArtifactId *string `json:"artifactId,omitempty"`
    // If set, auto-complete is enabled for this pull request and this is the identity that enabled it.
    AutoCompleteSetBy *IdentityRef `json:"autoCompleteSetBy,omitempty"`
    // The user who closed the pull request.
    ClosedBy *IdentityRef `json:"closedBy,omitempty"`
    // The date when the pull request was closed (completed, abandoned, or merged externally).
    ClosedDate *time.Time `json:"closedDate,omitempty"`
    // The code review ID of the pull request. Used internally.
    CodeReviewId *int `json:"codeReviewId,omitempty"`
    // The commits contained in the pull request.
    Commits *[]GitCommitRef `json:"commits,omitempty"`
    // Options which affect how the pull request will be merged when it is completed.
    CompletionOptions *GitPullRequestCompletionOptions `json:"completionOptions,omitempty"`
    // The most recent date at which the pull request entered the queue to be completed. Used internally.
    CompletionQueueTime *time.Time `json:"completionQueueTime,omitempty"`
    // The identity of the user who created the pull request.
    CreatedBy *IdentityRef `json:"createdBy,omitempty"`
    // The date when the pull request was created.
    CreationDate *time.Time `json:"creationDate,omitempty"`
    // The description of the pull request.
    Description *string `json:"description,omitempty"`
    // If this is a PR from a fork this will contain information about its source.
    ForkSource *GitForkRef `json:"forkSource,omitempty"`
    // Draft / WIP pull request.
    IsDraft *bool `json:"isDraft,omitempty"`
    // The labels associated with the pull request.
    Labels *[]WebApiTagDefinition `json:"labels,omitempty"`
    // The commit of the most recent pull request merge. If empty, the most recent merge is in progress or was unsuccessful.
    LastMergeCommit *GitCommitRef `json:"lastMergeCommit,omitempty"`
    // The commit at the head of the source branch at the time of the last pull request merge.
    LastMergeSourceCommit *GitCommitRef `json:"lastMergeSourceCommit,omitempty"`
    // The commit at the head of the target branch at the time of the last pull request merge.
    LastMergeTargetCommit *GitCommitRef `json:"lastMergeTargetCommit,omitempty"`
    // If set, pull request merge failed for this reason.
    MergeFailureMessage *string `json:"mergeFailureMessage,omitempty"`
    // The type of failure (if any) of the pull request merge.
    MergeFailureType *PullRequestMergeFailureType `json:"mergeFailureType,omitempty"`
    // The ID of the job used to run the pull request merge. Used internally.
    MergeId *uuid.UUID `json:"mergeId,omitempty"`
    // Options used when the pull request merge runs. These are separate from completion options since completion happens only once and a new merge will run every time the source branch of the pull request changes.
    MergeOptions *GitPullRequestMergeOptions `json:"mergeOptions,omitempty"`
    // The current status of the pull request merge.
    MergeStatus *PullRequestAsyncStatus `json:"mergeStatus,omitempty"`
    // The ID of the pull request.
    PullRequestId *int `json:"pullRequestId,omitempty"`
    // Used internally.
    RemoteUrl *string `json:"remoteUrl,omitempty"`
    // The repository containing the target branch of the pull request.
    Repository *GitRepository `json:"repository,omitempty"`
    // A list of reviewers on the pull request along with the state of their votes.
    Reviewers *[]IdentityRefWithVote `json:"reviewers,omitempty"`
    // The name of the source branch of the pull request.
    SourceRefName *string `json:"sourceRefName,omitempty"`
    // The status of the pull request.
    Status *PullRequestStatus `json:"status,omitempty"`
    // If true, this pull request supports multiple iterations. Iteration support means individual pushes to the source branch of the pull request can be reviewed and comments left in one iteration will be tracked across future iterations.
    SupportsIterations *bool `json:"supportsIterations,omitempty"`
    // The name of the target branch of the pull request.
    TargetRefName *string `json:"targetRefName,omitempty"`
    // The title of the pull request.
    Title *string `json:"title,omitempty"`
    // Used internally.
    Url *string `json:"url,omitempty"`
    // Any work item references associated with this pull request.
    WorkItemRefs *[]ResourceRef `json:"workItemRefs,omitempty"`
}

// Change made in a pull request.
type GitPullRequestChange struct {
    // ID used to track files through multiple changes.
    ChangeTrackingId *int `json:"changeTrackingId,omitempty"`
}

// Represents a comment thread of a pull request. A thread contains meta data about the file it was left on (if any) along with one or more comments (an initial comment and the subsequent replies).
type GitPullRequestCommentThread struct {
    // Links to other related objects.
    Links *ReferenceLinks `json:"_links,omitempty"`
    // A list of the comments.
    Comments *[]Comment `json:"comments,omitempty"`
    // The comment thread id.
    Id *int `json:"id,omitempty"`
    // Set of identities related to this thread
    Identities *map[string]IdentityRef `json:"identities,omitempty"`
    // Specify if the thread is deleted which happens when all comments are deleted.
    IsDeleted *bool `json:"isDeleted,omitempty"`
    // The time this thread was last updated.
    LastUpdatedDate *time.Time `json:"lastUpdatedDate,omitempty"`
    // Optional properties associated with the thread as a collection of key-value pairs.
    Properties *interface{} `json:"properties,omitempty"`
    // The time this thread was published.
    PublishedDate *time.Time `json:"publishedDate,omitempty"`
    // The status of the comment thread.
    Status *CommentThreadStatus `json:"status,omitempty"`
    // Specify thread context such as position in left/right file.
    ThreadContext *CommentThreadContext `json:"threadContext,omitempty"`
    // Extended context information unique to pull requests
    PullRequestThreadContext *GitPullRequestCommentThreadContext `json:"pullRequestThreadContext,omitempty"`
}

// Comment thread context contains details about what diffs were being viewed at the time of thread creation and whether or not the thread has been tracked from that original diff.
type GitPullRequestCommentThreadContext struct {
    // Used to track a comment across iterations. This value can be found by looking at the iteration's changes list. Must be set for pull requests with iteration support. Otherwise, it's not required for 'legacy' pull requests.
    ChangeTrackingId *int `json:"changeTrackingId,omitempty"`
    // The iteration context being viewed when the thread was created.
    IterationContext *CommentIterationContext `json:"iterationContext,omitempty"`
    // The criteria used to track this thread. If this property is filled out when the thread is returned, then the thread has been tracked from its original location using the given criteria.
    TrackingCriteria *CommentTrackingCriteria `json:"trackingCriteria,omitempty"`
}

// Preferences about how the pull request should be completed.
type GitPullRequestCompletionOptions struct {
    // If true, policies will be explicitly bypassed while the pull request is completed.
    BypassPolicy *bool `json:"bypassPolicy,omitempty"`
    // If policies are bypassed, this reason is stored as to why bypass was used.
    BypassReason *string `json:"bypassReason,omitempty"`
    // If true, the source branch of the pull request will be deleted after completion.
    DeleteSourceBranch *bool `json:"deleteSourceBranch,omitempty"`
    // If set, this will be used as the commit message of the merge commit.
    MergeCommitMessage *string `json:"mergeCommitMessage,omitempty"`
    // Specify the strategy used to merge the pull request during completion. If MergeStrategy is not set to any value, a no-FF merge will be created if SquashMerge == false. If MergeStrategy is not set to any value, the pull request commits will be squash if SquashMerge == true. The SquashMerge member is deprecated. It is recommended that you explicitly set MergeStrategy in all cases. If an explicit value is provided for MergeStrategy, the SquashMerge member will be ignored.
    MergeStrategy *GitPullRequestMergeStrategy `json:"mergeStrategy,omitempty"`
    // SquashMerge is deprecated. You should explicity set the value of MergeStrategy. If MergeStrategy is set to any value, the SquashMerge value will be ignored. If MergeStrategy is not set, the merge strategy will be no-fast-forward if this flag is false, or squash if true.
    SquashMerge *bool `json:"squashMerge,omitempty"`
    // If true, we will attempt to transition any work items linked to the pull request into the next logical state (i.e. Active -> Resolved)
    TransitionWorkItems *bool `json:"transitionWorkItems,omitempty"`
    // If true, the current completion attempt was triggered via auto-complete. Used internally.
    TriggeredByAutoComplete *bool `json:"triggeredByAutoComplete,omitempty"`
}

// Provides properties that describe a Git pull request iteration. Iterations are created as a result of creating and pushing updates to a pull request.
type GitPullRequestIteration struct {
    // A collection of related REST reference links.
    Links *ReferenceLinks `json:"_links,omitempty"`
    // Author of the pull request iteration.
    Author *IdentityRef `json:"author,omitempty"`
    // Changes included with the pull request iteration.
    ChangeList *[]GitPullRequestChange `json:"changeList,omitempty"`
    // The commits included with the pull request iteration.
    Commits *[]GitCommitRef `json:"commits,omitempty"`
    // The first common Git commit of the source and target refs.
    CommonRefCommit *GitCommitRef `json:"commonRefCommit,omitempty"`
    // The creation date of the pull request iteration.
    CreatedDate *time.Time `json:"createdDate,omitempty"`
    // Description of the pull request iteration.
    Description *string `json:"description,omitempty"`
    // Indicates if the Commits property contains a truncated list of commits in this pull request iteration.
    HasMoreCommits *bool `json:"hasMoreCommits,omitempty"`
    // ID of the pull request iteration. Iterations are created as a result of creating and pushing updates to a pull request.
    Id *int `json:"id,omitempty"`
    // If the iteration reason is Retarget, this is the refName of the new target
    NewTargetRefName *string `json:"newTargetRefName,omitempty"`
    // If the iteration reason is Retarget, this is the original target refName
    OldTargetRefName *string `json:"oldTargetRefName,omitempty"`
    // The Git push information associated with this pull request iteration.
    Push *GitPushRef `json:"push,omitempty"`
    // The reason for which the pull request iteration was created.
    Reason *IterationReason `json:"reason,omitempty"`
    // The source Git commit of this iteration.
    SourceRefCommit *GitCommitRef `json:"sourceRefCommit,omitempty"`
    // The target Git commit of this iteration.
    TargetRefCommit *GitCommitRef `json:"targetRefCommit,omitempty"`
    // The updated date of the pull request iteration.
    UpdatedDate *time.Time `json:"updatedDate,omitempty"`
}

// Collection of changes made in a pull request.
type GitPullRequestIterationChanges struct {
    // Changes made in the iteration.
    ChangeEntries *[]GitPullRequestChange `json:"changeEntries,omitempty"`
    // Value to specify as skip to get the next page of changes.  This will be zero if there are no more changes.
    NextSkip *int `json:"nextSkip,omitempty"`
    // Value to specify as top to get the next page of changes.  This will be zero if there are no more changes.
    NextTop *int `json:"nextTop,omitempty"`
}

// The options which are used when a pull request merge is created.
type GitPullRequestMergeOptions struct {
    DetectRenameFalsePositives *bool `json:"detectRenameFalsePositives,omitempty"`
    // If true, rename detection will not be performed during the merge.
    DisableRenames *bool `json:"disableRenames,omitempty"`
}

// Enumeration of possible merge strategies which can be used to complete a pull request.
type GitPullRequestMergeStrategy string

// A set of pull request queries and their results.
type GitPullRequestQuery struct {
    // The queries to perform.
    Queries *[]GitPullRequestQueryInput `json:"queries,omitempty"`
    // The results of the queries. This matches the QueryInputs list so Results[n] are the results of QueryInputs[n]. Each entry in the list is a dictionary of commit->pull requests.
    Results *[]map[string][]GitPullRequest `json:"results,omitempty"`
}

// Pull request query input parameters.
type GitPullRequestQueryInput struct {
    // The list of commit IDs to search for.
    Items *[]string `json:"items,omitempty"`
    // The type of query to perform.
    Type *GitPullRequestQueryType `json:"type,omitempty"`
}

// Accepted types of pull request queries.
type GitPullRequestQueryType string

type GitPullRequestReviewFileContentInfo struct {
    Links *ReferenceLinks `json:"_links,omitempty"`
    // The file change path.
    Path *string `json:"path,omitempty"`
    // Content hash of on-disk representation of file content. Its calculated by the client by using SHA1 hash function. Ensure that uploaded file has same encoding as in source control.
    SHA1Hash *string `json:"shA1Hash,omitempty"`
}

type GitPullRequestReviewFileType string

// Pull requests can be searched for matching this criteria.
type GitPullRequestSearchCriteria struct {
    // If set, search for pull requests that were created by this identity.
    CreatorId *uuid.UUID `json:"creatorId,omitempty"`
    // Whether to include the _links field on the shallow references
    IncludeLinks *bool `json:"includeLinks,omitempty"`
    // If set, search for pull requests whose target branch is in this repository.
    RepositoryId *uuid.UUID `json:"repositoryId,omitempty"`
    // If set, search for pull requests that have this identity as a reviewer.
    ReviewerId *uuid.UUID `json:"reviewerId,omitempty"`
    // If set, search for pull requests from this branch.
    SourceRefName *string `json:"sourceRefName,omitempty"`
    // If set, search for pull requests whose source branch is in this repository.
    SourceRepositoryId *uuid.UUID `json:"sourceRepositoryId,omitempty"`
    // If set, search for pull requests that are in this state. Defaults to Active if unset.
    Status *PullRequestStatus `json:"status,omitempty"`
    // If set, search for pull requests into this branch.
    TargetRefName *string `json:"targetRefName,omitempty"`
}

// This class contains the metadata of a service/extension posting pull request status. Status can be associated with a pull request or an iteration.
type GitPullRequestStatus struct {
    // Reference links.
    Links *ReferenceLinks `json:"_links,omitempty"`
    // Context of the status.
    Context *GitStatusContext `json:"context,omitempty"`
    // Identity that created the status.
    CreatedBy *IdentityRef `json:"createdBy,omitempty"`
    // Creation date and time of the status.
    CreationDate *time.Time `json:"creationDate,omitempty"`
    // Status description. Typically describes current state of the status.
    Description *string `json:"description,omitempty"`
    // Status identifier.
    Id *int `json:"id,omitempty"`
    // State of the status.
    State *GitStatusState `json:"state,omitempty"`
    // URL with status details.
    TargetUrl *string `json:"targetUrl,omitempty"`
    // Last update date and time of the status.
    UpdatedDate *time.Time `json:"updatedDate,omitempty"`
    // ID of the iteration to associate status with. Minimum value is 1.
    IterationId *int `json:"iterationId,omitempty"`
    // Custom properties of the status.
    Properties *interface{} `json:"properties,omitempty"`
}

type GitPush struct {
    Links *ReferenceLinks `json:"_links,omitempty"`
    Date *time.Time `json:"date,omitempty"`
    PushCorrelationId *uuid.UUID `json:"pushCorrelationId,omitempty"`
    PushedBy *IdentityRef `json:"pushedBy,omitempty"`
    PushId *int `json:"pushId,omitempty"`
    Url *string `json:"url,omitempty"`
    Commits *[]GitCommitRef `json:"commits,omitempty"`
    RefUpdates *[]GitRefUpdate `json:"refUpdates,omitempty"`
    Repository *GitRepository `json:"repository,omitempty"`
}

type GitPushEventData struct {
    AfterId *string `json:"afterId,omitempty"`
    BeforeId *string `json:"beforeId,omitempty"`
    Branch *string `json:"branch,omitempty"`
    Commits *[]GitCommit `json:"commits,omitempty"`
    Repository *GitRepository `json:"repository,omitempty"`
}

type GitPushRef struct {
    Links *ReferenceLinks `json:"_links,omitempty"`
    Date *time.Time `json:"date,omitempty"`
    // Deprecated: This is unused as of Dev15 M115 and may be deleted in the future
    PushCorrelationId *uuid.UUID `json:"pushCorrelationId,omitempty"`
    PushedBy *IdentityRef `json:"pushedBy,omitempty"`
    PushId *int `json:"pushId,omitempty"`
    Url *string `json:"url,omitempty"`
}

type GitPushSearchCriteria struct {
    FromDate *time.Time `json:"fromDate,omitempty"`
    // Whether to include the _links field on the shallow references
    IncludeLinks *bool `json:"includeLinks,omitempty"`
    IncludeRefUpdates *bool `json:"includeRefUpdates,omitempty"`
    PusherId *uuid.UUID `json:"pusherId,omitempty"`
    RefName *string `json:"refName,omitempty"`
    ToDate *time.Time `json:"toDate,omitempty"`
}

type GitQueryBranchStatsCriteria struct {
    BaseCommit *GitVersionDescriptor `json:"baseCommit,omitempty"`
    TargetCommits *[]GitVersionDescriptor `json:"targetCommits,omitempty"`
}

type GitQueryCommitsCriteria struct {
    // Number of entries to skip
    Skip *int `json:"$skip,omitempty"`
    // Maximum number of entries to retrieve
    Top *int `json:"$top,omitempty"`
    // Alias or display name of the author
    Author *string `json:"author,omitempty"`
    // Only applicable when ItemVersion specified. If provided, start walking history starting at this commit.
    CompareVersion *GitVersionDescriptor `json:"compareVersion,omitempty"`
    // Only applies when an itemPath is specified. This determines whether to exclude delete entries of the specified path.
    ExcludeDeletes *bool `json:"excludeDeletes,omitempty"`
    // If provided, a lower bound for filtering commits alphabetically
    FromCommitId *string `json:"fromCommitId,omitempty"`
    // If provided, only include history entries created after this date (string)
    FromDate *string `json:"fromDate,omitempty"`
    // What Git history mode should be used. This only applies to the search criteria when Ids = null and an itemPath is specified.
    HistoryMode *GitHistoryMode `json:"historyMode,omitempty"`
    // If provided, specifies the exact commit ids of the commits to fetch. May not be combined with other parameters.
    Ids *[]string `json:"ids,omitempty"`
    // Whether to include the _links field on the shallow references
    IncludeLinks *bool `json:"includeLinks,omitempty"`
    // Whether to include the push information
    IncludePushData *bool `json:"includePushData,omitempty"`
    // Whether to include the image Url for committers and authors
    IncludeUserImageUrl *bool `json:"includeUserImageUrl,omitempty"`
    // Whether to include linked work items
    IncludeWorkItems *bool `json:"includeWorkItems,omitempty"`
    // Path of item to search under
    ItemPath *string `json:"itemPath,omitempty"`
    // If provided, identifies the commit or branch to search
    ItemVersion *GitVersionDescriptor `json:"itemVersion,omitempty"`
    // If provided, an upper bound for filtering commits alphabetically
    ToCommitId *string `json:"toCommitId,omitempty"`
    // If provided, only include history entries created before this date (string)
    ToDate *string `json:"toDate,omitempty"`
    // Alias or display name of the committer
    User *string `json:"user,omitempty"`
}

type GitQueryRefsCriteria struct {
    // List of commit Ids to be searched
    CommitIds *[]string `json:"commitIds,omitempty"`
    // List of complete or partial names for refs to be searched
    RefNames *[]string `json:"refNames,omitempty"`
    // Type of search on refNames, if provided
    SearchType *GitRefSearchType `json:"searchType,omitempty"`
}

type GitRecycleBinRepositoryDetails struct {
    // Setting to false will undo earlier deletion and restore the repository.
    Deleted *bool `json:"deleted,omitempty"`
}

type GitRef struct {
    Links *ReferenceLinks `json:"_links,omitempty"`
    Creator *IdentityRef `json:"creator,omitempty"`
    IsLocked *bool `json:"isLocked,omitempty"`
    IsLockedBy *IdentityRef `json:"isLockedBy,omitempty"`
    Name *string `json:"name,omitempty"`
    ObjectId *string `json:"objectId,omitempty"`
    PeeledObjectId *string `json:"peeledObjectId,omitempty"`
    Statuses *[]GitStatus `json:"statuses,omitempty"`
    Url *string `json:"url,omitempty"`
}

type GitRefFavorite struct {
    Links *ReferenceLinks `json:"_links,omitempty"`
    Id *int `json:"id,omitempty"`
    IdentityId *uuid.UUID `json:"identityId,omitempty"`
    Name *string `json:"name,omitempty"`
    RepositoryId *uuid.UUID `json:"repositoryId,omitempty"`
    Type *RefFavoriteType `json:"type,omitempty"`
    Url *string `json:"url,omitempty"`
}

// Search type on ref name
type GitRefSearchType string

type GitRefUpdate struct {
    IsLocked *bool `json:"isLocked,omitempty"`
    Name *string `json:"name,omitempty"`
    NewObjectId *string `json:"newObjectId,omitempty"`
    OldObjectId *string `json:"oldObjectId,omitempty"`
    RepositoryId *uuid.UUID `json:"repositoryId,omitempty"`
}

// Enumerates the modes under which ref updates can be written to their repositories.
type GitRefUpdateMode string

type GitRefUpdateResult struct {
    // Custom message for the result object For instance, Reason for failing.
    CustomMessage *string `json:"customMessage,omitempty"`
    // Whether the ref is locked or not
    IsLocked *bool `json:"isLocked,omitempty"`
    // Ref name
    Name *string `json:"name,omitempty"`
    // New object ID
    NewObjectId *string `json:"newObjectId,omitempty"`
    // Old object ID
    OldObjectId *string `json:"oldObjectId,omitempty"`
    // Name of the plugin that rejected the updated.
    RejectedBy *string `json:"rejectedBy,omitempty"`
    // Repository ID
    RepositoryId *uuid.UUID `json:"repositoryId,omitempty"`
    // True if the ref update succeeded, false otherwise
    Success *bool `json:"success,omitempty"`
    // Status of the update from the TFS server.
    UpdateStatus *GitRefUpdateStatus `json:"updateStatus,omitempty"`
}

// Represents the possible outcomes from a request to update a ref in a repository.
type GitRefUpdateStatus string

type GitRepository struct {
    Links *ReferenceLinks `json:"_links,omitempty"`
    DefaultBranch *string `json:"defaultBranch,omitempty"`
    Id *uuid.UUID `json:"id,omitempty"`
    // True if the repository was created as a fork
    IsFork *bool `json:"isFork,omitempty"`
    Name *string `json:"name,omitempty"`
    ParentRepository *GitRepositoryRef `json:"parentRepository,omitempty"`
    Project *TeamProjectReference `json:"project,omitempty"`
    RemoteUrl *string `json:"remoteUrl,omitempty"`
    // Compressed size (bytes) of the repository.
    Size *uint64 `json:"size,omitempty"`
    SshUrl *string `json:"sshUrl,omitempty"`
    Url *string `json:"url,omitempty"`
    ValidRemoteUrls *[]string `json:"validRemoteUrls,omitempty"`
    WebUrl *string `json:"webUrl,omitempty"`
}

type GitRepositoryCreateOptions struct {
    Name *string `json:"name,omitempty"`
    ParentRepository *GitRepositoryRef `json:"parentRepository,omitempty"`
    Project *TeamProjectReference `json:"project,omitempty"`
}

type GitRepositoryRef struct {
    // Team Project Collection where this Fork resides
    Collection *TeamProjectCollectionReference `json:"collection,omitempty"`
    Id *uuid.UUID `json:"id,omitempty"`
    // True if the repository was created as a fork
    IsFork *bool `json:"isFork,omitempty"`
    Name *string `json:"name,omitempty"`
    Project *TeamProjectReference `json:"project,omitempty"`
    RemoteUrl *string `json:"remoteUrl,omitempty"`
    SshUrl *string `json:"sshUrl,omitempty"`
    Url *string `json:"url,omitempty"`
}

type GitRepositoryStats struct {
    ActivePullRequestsCount *int `json:"activePullRequestsCount,omitempty"`
    BranchesCount *int `json:"branchesCount,omitempty"`
    CommitsCount *int `json:"commitsCount,omitempty"`
    RepositoryId *string `json:"repositoryId,omitempty"`
}

type GitResolution struct {
    // User who created the resolution.
    Author *IdentityRef `json:"author,omitempty"`
}

// The type of a merge conflict.
type GitResolutionError string

type GitResolutionMergeContent struct {
    // User who created the resolution.
    Author *IdentityRef `json:"author,omitempty"`
    MergeType *GitResolutionMergeType `json:"mergeType,omitempty"`
    UserMergedBlob *GitBlobRef `json:"userMergedBlob,omitempty"`
    UserMergedContent *[]byte `json:"userMergedContent,omitempty"`
}

type GitResolutionMergeType string

type GitResolutionPathConflict struct {
    // User who created the resolution.
    Author *IdentityRef `json:"author,omitempty"`
    Action *GitResolutionPathConflictAction `json:"action,omitempty"`
    RenamePath *string `json:"renamePath,omitempty"`
}

type GitResolutionPathConflictAction string

type GitResolutionPickOneAction struct {
    // User who created the resolution.
    Author *IdentityRef `json:"author,omitempty"`
    Action *GitResolutionWhichAction `json:"action,omitempty"`
}

type GitResolutionRename1to2 struct {
    // User who created the resolution.
    Author *IdentityRef `json:"author,omitempty"`
    MergeType *GitResolutionMergeType `json:"mergeType,omitempty"`
    UserMergedBlob *GitBlobRef `json:"userMergedBlob,omitempty"`
    UserMergedContent *[]byte `json:"userMergedContent,omitempty"`
    Action *GitResolutionRename1to2Action `json:"action,omitempty"`
}

type GitResolutionRename1to2Action string

// Resolution status of a conflict.
type GitResolutionStatus string

type GitResolutionWhichAction string

type GitRevert struct {
    Links *ReferenceLinks `json:"_links,omitempty"`
    DetailedStatus *GitAsyncRefOperationDetail `json:"detailedStatus,omitempty"`
    Parameters *GitAsyncRefOperationParameters `json:"parameters,omitempty"`
    Status *GitAsyncOperationStatus `json:"status,omitempty"`
    // A URL that can be used to make further requests for status about the operation
    Url *string `json:"url,omitempty"`
    RevertId *int `json:"revertId,omitempty"`
}

// This class contains the metadata of a service/extension posting a status.
type GitStatus struct {
    // Reference links.
    Links *ReferenceLinks `json:"_links,omitempty"`
    // Context of the status.
    Context *GitStatusContext `json:"context,omitempty"`
    // Identity that created the status.
    CreatedBy *IdentityRef `json:"createdBy,omitempty"`
    // Creation date and time of the status.
    CreationDate *time.Time `json:"creationDate,omitempty"`
    // Status description. Typically describes current state of the status.
    Description *string `json:"description,omitempty"`
    // Status identifier.
    Id *int `json:"id,omitempty"`
    // State of the status.
    State *GitStatusState `json:"state,omitempty"`
    // URL with status details.
    TargetUrl *string `json:"targetUrl,omitempty"`
    // Last update date and time of the status.
    UpdatedDate *time.Time `json:"updatedDate,omitempty"`
}

// Status context that uniquely identifies the status.
type GitStatusContext struct {
    // Genre of the status. Typically name of the service/tool generating the status, can be empty.
    Genre *string `json:"genre,omitempty"`
    // Name identifier of the status, cannot be null or empty.
    Name *string `json:"name,omitempty"`
}

// State of the status.
type GitStatusState string

// An object describing the git suggestion.  Git suggestions are currently limited to suggested pull requests.
type GitSuggestion struct {
    // Specific properties describing the suggestion.
    Properties *map[string]interface{} `json:"properties,omitempty"`
    // The type of suggestion (e.g. pull request).
    Type *string `json:"type,omitempty"`
}

type GitTargetVersionDescriptor struct {
    // Version string identifier (name of tag/branch, SHA1 of commit)
    Version *string `json:"version,omitempty"`
    // Version options - Specify additional modifiers to version (e.g Previous)
    VersionOptions *GitVersionOptions `json:"versionOptions,omitempty"`
    // Version type (branch, tag, or commit). Determines how Id is interpreted
    VersionType *GitVersionType `json:"versionType,omitempty"`
    // Version string identifier (name of tag/branch, SHA1 of commit)
    TargetVersion *string `json:"targetVersion,omitempty"`
    // Version options - Specify additional modifiers to version (e.g Previous)
    TargetVersionOptions *GitVersionOptions `json:"targetVersionOptions,omitempty"`
    // Version type (branch, tag, or commit). Determines how Id is interpreted
    TargetVersionType *GitVersionType `json:"targetVersionType,omitempty"`
}

type GitTemplate struct {
    // Name of the Template
    Name *string `json:"name,omitempty"`
    // Type of the Template
    Type *string `json:"type,omitempty"`
}

type GitTreeDiff struct {
    // ObjectId of the base tree of this diff.
    BaseTreeId *string `json:"baseTreeId,omitempty"`
    // List of tree entries that differ between the base and target tree.  Renames and object type changes are returned as a delete for the old object and add for the new object.  If a continuation token is returned in the response header, some tree entries are yet to be processed and may yeild more diff entries. If the continuation token is not returned all the diff entries have been included in this response.
    DiffEntries *[]GitTreeDiffEntry `json:"diffEntries,omitempty"`
    // ObjectId of the target tree of this diff.
    TargetTreeId *string `json:"targetTreeId,omitempty"`
    // REST Url to this resource.
    Url *string `json:"url,omitempty"`
}

type GitTreeDiffEntry struct {
    // SHA1 hash of the object in the base tree, if it exists. Will be null in case of adds.
    BaseObjectId *string `json:"baseObjectId,omitempty"`
    // Type of change that affected this entry.
    ChangeType *VersionControlChangeType `json:"changeType,omitempty"`
    // Object type of the tree entry. Blob, Tree or Commit("submodule")
    ObjectType *GitObjectType `json:"objectType,omitempty"`
    // Relative path in base and target trees.
    Path *string `json:"path,omitempty"`
    // SHA1 hash of the object in the target tree, if it exists. Will be null in case of deletes.
    TargetObjectId *string `json:"targetObjectId,omitempty"`
}

type GitTreeDiffResponse struct {
    // The HTTP client methods find the continuation token header in the response and populate this field.
    ContinuationToken *[]string `json:"continuationToken,omitempty"`
    TreeDiff *GitTreeDiff `json:"treeDiff,omitempty"`
}

type GitTreeEntryRef struct {
    // Blob or tree
    GitObjectType *GitObjectType `json:"gitObjectType,omitempty"`
    // Mode represented as octal string
    Mode *string `json:"mode,omitempty"`
    // SHA1 hash of git object
    ObjectId *string `json:"objectId,omitempty"`
    // Path relative to parent tree object
    RelativePath *string `json:"relativePath,omitempty"`
    // Size of content
    Size *uint64 `json:"size,omitempty"`
    // url to retrieve tree or blob
    Url *string `json:"url,omitempty"`
}

type GitTreeRef struct {
    Links *ReferenceLinks `json:"_links,omitempty"`
    // SHA1 hash of git object
    ObjectId *string `json:"objectId,omitempty"`
    // Sum of sizes of all children
    Size *uint64 `json:"size,omitempty"`
    // Blobs and trees under this tree
    TreeEntries *[]GitTreeEntryRef `json:"treeEntries,omitempty"`
    // Url to tree
    Url *string `json:"url,omitempty"`
}

// User info and date for Git operations.
type GitUserDate struct {
    // Date of the Git operation.
    Date *time.Time `json:"date,omitempty"`
    // Email address of the user performing the Git operation.
    Email *string `json:"email,omitempty"`
    // Url for the user's avatar.
    ImageUrl *string `json:"imageUrl,omitempty"`
    // Name of the user performing the Git operation.
    Name *string `json:"name,omitempty"`
}

type GitVersionDescriptor struct {
    // Version string identifier (name of tag/branch, SHA1 of commit)
    Version *string `json:"version,omitempty"`
    // Version options - Specify additional modifiers to version (e.g Previous)
    VersionOptions *GitVersionOptions `json:"versionOptions,omitempty"`
    // Version type (branch, tag, or commit). Determines how Id is interpreted
    VersionType *GitVersionType `json:"versionType,omitempty"`
}

// Accepted types of version options
type GitVersionOptions string

// Accepted types of version
type GitVersionType string

// Globally unique key for a repository.
type GlobalGitRepositoryKey struct {
    // Team Project Collection ID of the collection for the repository.
    CollectionId *uuid.UUID `json:"collectionId,omitempty"`
    // Team Project ID of the project for the repository.
    ProjectId *uuid.UUID `json:"projectId,omitempty"`
    // ID of the repository.
    RepositoryId *uuid.UUID `json:"repositoryId,omitempty"`
}

type GraphSubjectBase struct {
    // This field contains zero or more interesting links about the graph subject. These links may be invoked to obtain additional relationships or more detailed information about this graph subject.
    Links *ReferenceLinks `json:"_links,omitempty"`
    // The descriptor is the primary way to reference the graph subject while the system is running. This field will uniquely identify the same graph subject across both Accounts and Organizations.
    Descriptor *string `json:"descriptor,omitempty"`
    // This is the non-unique display name of the graph subject. To change this field, you must alter its value in the source provider.
    DisplayName *string `json:"displayName,omitempty"`
    // This url is the full route to the source resource of this graph subject.
    Url *string `json:"url,omitempty"`
}

type HistoryEntry struct {
    // The Change list (changeset/commit/shelveset) for this point in history
    ChangeList *ChangeList `json:"changeList,omitempty"`
    // The change made to the item from this change list (only relevant for File history, not folders)
    ItemChangeType *VersionControlChangeType `json:"itemChangeType,omitempty"`
    // The path of the item at this point in history (only relevant for File history, not folders)
    ServerItem *string `json:"serverItem,omitempty"`
}

type IdentityRef struct {
    // This field contains zero or more interesting links about the graph subject. These links may be invoked to obtain additional relationships or more detailed information about this graph subject.
    Links *ReferenceLinks `json:"_links,omitempty"`
    // The descriptor is the primary way to reference the graph subject while the system is running. This field will uniquely identify the same graph subject across both Accounts and Organizations.
    Descriptor *string `json:"descriptor,omitempty"`
    // This is the non-unique display name of the graph subject. To change this field, you must alter its value in the source provider.
    DisplayName *string `json:"displayName,omitempty"`
    // This url is the full route to the source resource of this graph subject.
    Url *string `json:"url,omitempty"`
    // Deprecated - Can be retrieved by querying the Graph user referenced in the "self" entry of the IdentityRef "_links" dictionary
    DirectoryAlias *string `json:"directoryAlias,omitempty"`
    Id *string `json:"id,omitempty"`
    // Deprecated - Available in the "avatar" entry of the IdentityRef "_links" dictionary
    ImageUrl *string `json:"imageUrl,omitempty"`
    // Deprecated - Can be retrieved by querying the Graph membership state referenced in the "membershipState" entry of the GraphUser "_links" dictionary
    Inactive *bool `json:"inactive,omitempty"`
    // Deprecated - Can be inferred from the subject type of the descriptor (Descriptor.IsAadUserType/Descriptor.IsAadGroupType)
    IsAadIdentity *bool `json:"isAadIdentity,omitempty"`
    // Deprecated - Can be inferred from the subject type of the descriptor (Descriptor.IsGroupType)
    IsContainer *bool `json:"isContainer,omitempty"`
    IsDeletedInOrigin *bool `json:"isDeletedInOrigin,omitempty"`
    // Deprecated - not in use in most preexisting implementations of ToIdentityRef
    ProfileUrl *string `json:"profileUrl,omitempty"`
    // Deprecated - use Domain+PrincipalName instead
    UniqueName *string `json:"uniqueName,omitempty"`
}

// Identity information including a vote on a pull request.
type IdentityRefWithVote struct {
    // This field contains zero or more interesting links about the graph subject. These links may be invoked to obtain additional relationships or more detailed information about this graph subject.
    Links *ReferenceLinks `json:"_links,omitempty"`
    // The descriptor is the primary way to reference the graph subject while the system is running. This field will uniquely identify the same graph subject across both Accounts and Organizations.
    Descriptor *string `json:"descriptor,omitempty"`
    // This is the non-unique display name of the graph subject. To change this field, you must alter its value in the source provider.
    DisplayName *string `json:"displayName,omitempty"`
    // This url is the full route to the source resource of this graph subject.
    Url *string `json:"url,omitempty"`
    // Deprecated - Can be retrieved by querying the Graph user referenced in the "self" entry of the IdentityRef "_links" dictionary
    DirectoryAlias *string `json:"directoryAlias,omitempty"`
    Id *string `json:"id,omitempty"`
    // Deprecated - Available in the "avatar" entry of the IdentityRef "_links" dictionary
    ImageUrl *string `json:"imageUrl,omitempty"`
    // Deprecated - Can be retrieved by querying the Graph membership state referenced in the "membershipState" entry of the GraphUser "_links" dictionary
    Inactive *bool `json:"inactive,omitempty"`
    // Deprecated - Can be inferred from the subject type of the descriptor (Descriptor.IsAadUserType/Descriptor.IsAadGroupType)
    IsAadIdentity *bool `json:"isAadIdentity,omitempty"`
    // Deprecated - Can be inferred from the subject type of the descriptor (Descriptor.IsGroupType)
    IsContainer *bool `json:"isContainer,omitempty"`
    IsDeletedInOrigin *bool `json:"isDeletedInOrigin,omitempty"`
    // Deprecated - not in use in most preexisting implementations of ToIdentityRef
    ProfileUrl *string `json:"profileUrl,omitempty"`
    // Deprecated - use Domain+PrincipalName instead
    UniqueName *string `json:"uniqueName,omitempty"`
    // Indicates if this is a required reviewer for this pull request. <br /> Branches can have policies that require particular reviewers are required for pull requests.
    IsRequired *bool `json:"isRequired,omitempty"`
    // URL to retrieve information about this identity
    ReviewerUrl *string `json:"reviewerUrl,omitempty"`
    // Vote on a pull request:<br /> 10 - approved 5 - approved with suggestions 0 - no vote -5 - waiting for author -10 - rejected
    Vote *int `json:"vote,omitempty"`
    // Groups or teams that that this reviewer contributed to. <br /> Groups and teams can be reviewers on pull requests but can not vote directly.  When a member of the group or team votes, that vote is rolled up into the group or team vote.  VotedFor is a list of such votes.
    VotedFor *[]IdentityRefWithVote `json:"votedFor,omitempty"`
}

type ImportRepositoryValidation struct {
    GitSource *GitImportGitSource `json:"gitSource,omitempty"`
    Password *string `json:"password,omitempty"`
    TfvcSource *GitImportTfvcSource `json:"tfvcSource,omitempty"`
    Username *string `json:"username,omitempty"`
}

type IncludedGitCommit struct {
    CommitId *string `json:"commitId,omitempty"`
    CommitTime *time.Time `json:"commitTime,omitempty"`
    ParentCommitIds *[]string `json:"parentCommitIds,omitempty"`
    RepositoryId *uuid.UUID `json:"repositoryId,omitempty"`
}

// Real time event (SignalR) for IsDraft update on a pull request
type IsDraftUpdatedEvent struct {
    // The id of this event. Can be used to track send/receive state between client and server.
    EventId *uuid.UUID `json:"eventId,omitempty"`
    // The id of the pull request this event was generated for.
    PullRequestId *int `json:"pullRequestId,omitempty"`
}

type ItemContent struct {
    Content *string `json:"content,omitempty"`
    ContentType *ItemContentType `json:"contentType,omitempty"`
}

type ItemContentType string

// Optional details to include when returning an item model
type ItemDetailsOptions struct {
    // If true, include metadata about the file type
    IncludeContentMetadata *bool `json:"includeContentMetadata,omitempty"`
    // Specifies whether to include children (OneLevel), all descendants (Full) or None for folder items
    RecursionLevel *VersionControlRecursionType `json:"recursionLevel,omitempty"`
}

type ItemModel struct {
    Links *ReferenceLinks `json:"_links,omitempty"`
    Content *string `json:"content,omitempty"`
    ContentMetadata *FileContentMetadata `json:"contentMetadata,omitempty"`
    IsFolder *bool `json:"isFolder,omitempty"`
    IsSymLink *bool `json:"isSymLink,omitempty"`
    Path *string `json:"path,omitempty"`
    Url *string `json:"url,omitempty"`
}

// The reason for which the pull request iteration was created.
type IterationReason string

// The JSON model for a JSON Patch operation
type JsonPatchOperation struct {
    // The path to copy from for the Move/Copy operation.
    From *string `json:"from,omitempty"`
    // The patch operation
    Op *Operation `json:"op,omitempty"`
    // The path for the operation. In the case of an array, a zero based index can be used to specify the position in the array (e.g. /biscuits/0/name). The "-" character can be used instead of an index to insert at the end of the array (e.g. /biscuits/-).
    Path *string `json:"path,omitempty"`
    // The value for the operation. This is either a primitive or a JToken.
    Value *interface{} `json:"value,omitempty"`
}

// Real time event (SignalR) for updated labels on a pull request
type LabelsUpdatedEvent struct {
    // The id of this event. Can be used to track send/receive state between client and server.
    EventId *uuid.UUID `json:"eventId,omitempty"`
    // The id of the pull request this event was generated for.
    PullRequestId *int `json:"pullRequestId,omitempty"`
}

// The class to represent the line diff block
type LineDiffBlock struct {
    // Type of change that was made to the block.
    ChangeType *LineDiffBlockChangeType `json:"changeType,omitempty"`
    // Line number where this block starts in modified file.
    ModifiedLineNumberStart *int `json:"modifiedLineNumberStart,omitempty"`
    // Count of lines in this block in modified file.
    ModifiedLinesCount *int `json:"modifiedLinesCount,omitempty"`
    // Line number where this block starts in original file.
    OriginalLineNumberStart *int `json:"originalLineNumberStart,omitempty"`
    // Count of lines in this block in original file.
    OriginalLinesCount *int `json:"originalLinesCount,omitempty"`
}

// Type of change for a line diff block
type LineDiffBlockChangeType string

// Real time event (SignalR) for a merge completed on a pull request
type MergeCompletedEvent struct {
    // The id of this event. Can be used to track send/receive state between client and server.
    EventId *uuid.UUID `json:"eventId,omitempty"`
    // The id of the pull request this event was generated for.
    PullRequestId *int `json:"pullRequestId,omitempty"`
}

type Operation string

// The full policy configuration with settings.
type PolicyConfiguration struct {
    // The policy configuration ID.
    Id *int `json:"id,omitempty"`
    // The policy configuration type.
    Type *PolicyTypeRef `json:"type,omitempty"`
    // The URL where the policy configuration can be retrieved.
    Url *string `json:"url,omitempty"`
    // The policy configuration revision ID.
    Revision *int `json:"revision,omitempty"`
    // The links to other objects related to this object.
    Links *ReferenceLinks `json:"_links,omitempty"`
    // A reference to the identity that created the policy.
    CreatedBy *IdentityRef `json:"createdBy,omitempty"`
    // The date and time when the policy was created.
    CreatedDate *time.Time `json:"createdDate,omitempty"`
    // Indicates whether the policy is blocking.
    IsBlocking *bool `json:"isBlocking,omitempty"`
    // Indicates whether the policy has been (soft) deleted.
    IsDeleted *bool `json:"isDeleted,omitempty"`
    // Indicates whether the policy is enabled.
    IsEnabled *bool `json:"isEnabled,omitempty"`
    // The policy configuration settings.
    Settings *interface{} `json:"settings,omitempty"`
}

// Policy configuration reference.
type PolicyConfigurationRef struct {
    // The policy configuration ID.
    Id *int `json:"id,omitempty"`
    // The policy configuration type.
    Type *PolicyTypeRef `json:"type,omitempty"`
    // The URL where the policy configuration can be retrieved.
    Url *string `json:"url,omitempty"`
}

// Real time event (SignalR) for a policy evaluation update on a pull request
type PolicyEvaluationUpdatedEvent struct {
    // The id of this event. Can be used to track send/receive state between client and server.
    EventId *uuid.UUID `json:"eventId,omitempty"`
    // The id of the pull request this event was generated for.
    PullRequestId *int `json:"pullRequestId,omitempty"`
}

// Policy type reference.
type PolicyTypeRef struct {
    // Display name of the policy type.
    DisplayName *string `json:"displayName,omitempty"`
    // The policy type ID.
    Id *uuid.UUID `json:"id,omitempty"`
    // The URL where the policy type can be retrieved.
    Url *string `json:"url,omitempty"`
}

type ProjectVisibility string

// The status of a pull request merge.
type PullRequestAsyncStatus string

// Real time event (SignalR) for pull request creation
type PullRequestCreatedEvent struct {
    // The id of this event. Can be used to track send/receive state between client and server.
    EventId *uuid.UUID `json:"eventId,omitempty"`
    // The id of the pull request this event was generated for.
    PullRequestId *int `json:"pullRequestId,omitempty"`
}

// The specific type of a pull request merge failure.
type PullRequestMergeFailureType string

// Status of a pull request.
type PullRequestStatus string

// Initial config contract sent to extensions creating tabs on the pull request page
type PullRequestTabExtensionConfig struct {
    PullRequestId *int `json:"pullRequestId,omitempty"`
    RepositoryId *string `json:"repositoryId,omitempty"`
}

// Base contract for a real time pull request event (SignalR)
type RealTimePullRequestEvent struct {
    // The id of this event. Can be used to track send/receive state between client and server.
    EventId *uuid.UUID `json:"eventId,omitempty"`
    // The id of the pull request this event was generated for.
    PullRequestId *int `json:"pullRequestId,omitempty"`
}

// The class to represent a collection of REST reference links.
type ReferenceLinks struct {
    // The readonly view of the links.  Because Reference links are readonly, we only want to expose them as read only.
    Links *map[string]interface{} `json:"links,omitempty"`
}

type RefFavoriteType string

type ResourceRef struct {
    Id *string `json:"id,omitempty"`
    Url *string `json:"url,omitempty"`
}

// Real time event (SignalR) for when the target branch of a pull request is changed
type RetargetEvent struct {
    // The id of this event. Can be used to track send/receive state between client and server.
    EventId *uuid.UUID `json:"eventId,omitempty"`
    // The id of the pull request this event was generated for.
    PullRequestId *int `json:"pullRequestId,omitempty"`
}

// Real time event (SignalR) for an update to reviewers on a pull request
type ReviewersUpdatedEvent struct {
    // The id of this event. Can be used to track send/receive state between client and server.
    EventId *uuid.UUID `json:"eventId,omitempty"`
    // The id of the pull request this event was generated for.
    PullRequestId *int `json:"pullRequestId,omitempty"`
}

// Real time event (SignalR) for reviewer votes being reset on a pull request
type ReviewersVotesResetEvent struct {
    // The id of this event. Can be used to track send/receive state between client and server.
    EventId *uuid.UUID `json:"eventId,omitempty"`
    // The id of the pull request this event was generated for.
    PullRequestId *int `json:"pullRequestId,omitempty"`
}

// Real time event (SignalR) for a reviewer vote update on a pull request
type ReviewerVoteUpdatedEvent struct {
    // The id of this event. Can be used to track send/receive state between client and server.
    EventId *uuid.UUID `json:"eventId,omitempty"`
    // The id of the pull request this event was generated for.
    PullRequestId *int `json:"pullRequestId,omitempty"`
}

// Context used while sharing a pull request.
type ShareNotificationContext struct {
    // Optional user note or message.
    Message *string `json:"message,omitempty"`
    // Identities of users who will receive a share notification.
    Receivers *[]IdentityRef `json:"receivers,omitempty"`
}

type SourceControlTypes string

type SourceToTargetRef struct {
    // The source ref to copy. For example, refs/heads/master.
    SourceRef *string `json:"sourceRef,omitempty"`
    // The target ref to update. For example, refs/heads/master.
    TargetRef *string `json:"targetRef,omitempty"`
}

// Real time event (SignalR) for an added status on a pull request
type StatusAddedEvent struct {
    // The id of this event. Can be used to track send/receive state between client and server.
    EventId *uuid.UUID `json:"eventId,omitempty"`
    // The id of the pull request this event was generated for.
    PullRequestId *int `json:"pullRequestId,omitempty"`
}

// Real time event (SignalR) for deleted statuses on a pull request
type StatusesDeletedEvent struct {
    // The id of this event. Can be used to track send/receive state between client and server.
    EventId *uuid.UUID `json:"eventId,omitempty"`
    // The id of the pull request this event was generated for.
    PullRequestId *int `json:"pullRequestId,omitempty"`
}

// Real time event (SignalR) for a status update on a pull request
type StatusUpdatedEvent struct {
    // The id of this event. Can be used to track send/receive state between client and server.
    EventId *uuid.UUID `json:"eventId,omitempty"`
    // The id of the pull request this event was generated for.
    PullRequestId *int `json:"pullRequestId,omitempty"`
}

// Represents a Supported IDE entity.
type SupportedIde struct {
    // The download URL for the IDE.
    DownloadUrl *string `json:"downloadUrl,omitempty"`
    // The type of the IDE.
    IdeType *SupportedIdeType `json:"ideType,omitempty"`
    // The name of the IDE.
    Name *string `json:"name,omitempty"`
    // The URL to open the protocol handler for the IDE.
    ProtocolHandlerUrl *string `json:"protocolHandlerUrl,omitempty"`
    // A list of SupportedPlatforms.
    SupportedPlatforms *[]string `json:"supportedPlatforms,omitempty"`
}

// Enumeration that represents the types of IDEs supported.
type SupportedIdeType string

// Reference object for a TeamProjectCollection.
type TeamProjectCollectionReference struct {
    // Collection Id.
    Id *uuid.UUID `json:"id,omitempty"`
    // Collection Name.
    Name *string `json:"name,omitempty"`
    // Collection REST Url.
    Url *string `json:"url,omitempty"`
}

// Represents a shallow reference to a TeamProject.
type TeamProjectReference struct {
    // Project abbreviation.
    Abbreviation *string `json:"abbreviation,omitempty"`
    // Url to default team identity image.
    DefaultTeamImageUrl *string `json:"defaultTeamImageUrl,omitempty"`
    // The project's description (if any).
    Description *string `json:"description,omitempty"`
    // Project identifier.
    Id *uuid.UUID `json:"id,omitempty"`
    // Project last update time.
    LastUpdateTime *time.Time `json:"lastUpdateTime,omitempty"`
    // Project name.
    Name *string `json:"name,omitempty"`
    // Project revision.
    Revision *uint64 `json:"revision,omitempty"`
    // Project state.
    State *interface{} `json:"state,omitempty"`
    // Url to the full version of the object.
    Url *string `json:"url,omitempty"`
    // Project visibility.
    Visibility *ProjectVisibility `json:"visibility,omitempty"`
}

type TfvcBranch struct {
    // Path for the branch.
    Path *string `json:"path,omitempty"`
    // A collection of REST reference links.
    Links *ReferenceLinks `json:"_links,omitempty"`
    // Creation date of the branch.
    CreatedDate *time.Time `json:"createdDate,omitempty"`
    // Description of the branch.
    Description *string `json:"description,omitempty"`
    // Is the branch deleted?
    IsDeleted *bool `json:"isDeleted,omitempty"`
    // Alias or display name of user
    Owner *IdentityRef `json:"owner,omitempty"`
    // URL to retrieve the item.
    Url *string `json:"url,omitempty"`
    // List of children for the branch.
    Children *[]TfvcBranch `json:"children,omitempty"`
    // List of branch mappings.
    Mappings *[]TfvcBranchMapping `json:"mappings,omitempty"`
    // Path of the branch's parent.
    Parent *TfvcShallowBranchRef `json:"parent,omitempty"`
    // List of paths of the related branches.
    RelatedBranches *[]TfvcShallowBranchRef `json:"relatedBranches,omitempty"`
}

type TfvcBranchMapping struct {
    // Depth of the branch.
    Depth *string `json:"depth,omitempty"`
    // Server item for the branch.
    ServerItem *string `json:"serverItem,omitempty"`
    // Type of the branch.
    Type *string `json:"type,omitempty"`
}

type TfvcBranchRef struct {
    // Path for the branch.
    Path *string `json:"path,omitempty"`
    // A collection of REST reference links.
    Links *ReferenceLinks `json:"_links,omitempty"`
    // Creation date of the branch.
    CreatedDate *time.Time `json:"createdDate,omitempty"`
    // Description of the branch.
    Description *string `json:"description,omitempty"`
    // Is the branch deleted?
    IsDeleted *bool `json:"isDeleted,omitempty"`
    // Alias or display name of user
    Owner *IdentityRef `json:"owner,omitempty"`
    // URL to retrieve the item.
    Url *string `json:"url,omitempty"`
}

type TfvcChange struct {
    // List of merge sources in case of rename or branch creation.
    MergeSources *[]TfvcMergeSource `json:"mergeSources,omitempty"`
    // Version at which a (shelved) change was pended against
    PendingVersion *int `json:"pendingVersion,omitempty"`
}

type TfvcChangeset struct {
    // A collection of REST reference links.
    Links *ReferenceLinks `json:"_links,omitempty"`
    // Alias or display name of user
    Author *IdentityRef `json:"author,omitempty"`
    // Id of the changeset.
    ChangesetId *int `json:"changesetId,omitempty"`
    // Alias or display name of user
    CheckedInBy *IdentityRef `json:"checkedInBy,omitempty"`
    // Comment for the changeset.
    Comment *string `json:"comment,omitempty"`
    // Was the Comment result truncated?
    CommentTruncated *bool `json:"commentTruncated,omitempty"`
    // Creation date of the changeset.
    CreatedDate *time.Time `json:"createdDate,omitempty"`
    // URL to retrieve the item.
    Url *string `json:"url,omitempty"`
    // Account Id of the changeset.
    AccountId *uuid.UUID `json:"accountId,omitempty"`
    // List of associated changes.
    Changes *[]TfvcChange `json:"changes,omitempty"`
    // Checkin Notes for the changeset.
    CheckinNotes *[]CheckinNote `json:"checkinNotes,omitempty"`
    // Collection Id of the changeset.
    CollectionId *uuid.UUID `json:"collectionId,omitempty"`
    // Are more changes available.
    HasMoreChanges *bool `json:"hasMoreChanges,omitempty"`
    // Policy Override for the changeset.
    PolicyOverride *TfvcPolicyOverrideInfo `json:"policyOverride,omitempty"`
    // Team Project Ids for the changeset.
    TeamProjectIds *[]uuid.UUID `json:"teamProjectIds,omitempty"`
    // List of work items associated with the changeset.
    WorkItems *[]AssociatedWorkItem `json:"workItems,omitempty"`
}

type TfvcChangesetRef struct {
    // A collection of REST reference links.
    Links *ReferenceLinks `json:"_links,omitempty"`
    // Alias or display name of user
    Author *IdentityRef `json:"author,omitempty"`
    // Id of the changeset.
    ChangesetId *int `json:"changesetId,omitempty"`
    // Alias or display name of user
    CheckedInBy *IdentityRef `json:"checkedInBy,omitempty"`
    // Comment for the changeset.
    Comment *string `json:"comment,omitempty"`
    // Was the Comment result truncated?
    CommentTruncated *bool `json:"commentTruncated,omitempty"`
    // Creation date of the changeset.
    CreatedDate *time.Time `json:"createdDate,omitempty"`
    // URL to retrieve the item.
    Url *string `json:"url,omitempty"`
}

// Criteria used in a search for change lists
type TfvcChangesetSearchCriteria struct {
    // Alias or display name of user who made the changes
    Author *string `json:"author,omitempty"`
    // Whether or not to follow renames for the given item being queried
    FollowRenames *bool `json:"followRenames,omitempty"`
    // If provided, only include changesets created after this date (string) Think of a better name for this.
    FromDate *string `json:"fromDate,omitempty"`
    // If provided, only include changesets after this changesetID
    FromId *int `json:"fromId,omitempty"`
    // Whether to include the _links field on the shallow references
    IncludeLinks *bool `json:"includeLinks,omitempty"`
    // Path of item to search under
    ItemPath *string `json:"itemPath,omitempty"`
    Mappings *[]TfvcMappingFilter `json:"mappings,omitempty"`
    // If provided, only include changesets created before this date (string) Think of a better name for this.
    ToDate *string `json:"toDate,omitempty"`
    // If provided, a version descriptor for the latest change list to include
    ToId *int `json:"toId,omitempty"`
}

type TfvcChangesetsRequestData struct {
    // List of changeset Ids.
    ChangesetIds *[]int `json:"changesetIds,omitempty"`
    // Length of the comment.
    CommentLength *int `json:"commentLength,omitempty"`
    // Whether to include the _links field on the shallow references
    IncludeLinks *bool `json:"includeLinks,omitempty"`
}

type TfvcCheckinEventData struct {
    Changeset *TfvcChangeset `json:"changeset,omitempty"`
    Project *TeamProjectReference `json:"project,omitempty"`
}

type TfvcHistoryEntry struct {
    // The encoding of the item at this point in history (only relevant for File history, not folders)
    Encoding *int `json:"encoding,omitempty"`
    // The file id of the item at this point in history (only relevant for File history, not folders)
    FileId *int `json:"fileId,omitempty"`
}

type TfvcItem struct {
    Links *ReferenceLinks `json:"_links,omitempty"`
    Content *string `json:"content,omitempty"`
    ContentMetadata *FileContentMetadata `json:"contentMetadata,omitempty"`
    IsFolder *bool `json:"isFolder,omitempty"`
    IsSymLink *bool `json:"isSymLink,omitempty"`
    Path *string `json:"path,omitempty"`
    Url *string `json:"url,omitempty"`
    ChangeDate *time.Time `json:"changeDate,omitempty"`
    DeletionId *int `json:"deletionId,omitempty"`
    // File encoding from database, -1 represents binary.
    Encoding *int `json:"encoding,omitempty"`
    // MD5 hash as a base 64 string, applies to files only.
    HashValue *string `json:"hashValue,omitempty"`
    IsBranch *bool `json:"isBranch,omitempty"`
    IsPendingChange *bool `json:"isPendingChange,omitempty"`
    // The size of the file, if applicable.
    Size *uint64 `json:"size,omitempty"`
    Version *int `json:"version,omitempty"`
}

// Item path and Version descriptor properties
type TfvcItemDescriptor struct {
    Path *string `json:"path,omitempty"`
    RecursionLevel *VersionControlRecursionType `json:"recursionLevel,omitempty"`
    Version *string `json:"version,omitempty"`
    VersionOption *TfvcVersionOption `json:"versionOption,omitempty"`
    VersionType *TfvcVersionType `json:"versionType,omitempty"`
}

type TfvcItemPreviousHash struct {
    Links *ReferenceLinks `json:"_links,omitempty"`
    Content *string `json:"content,omitempty"`
    ContentMetadata *FileContentMetadata `json:"contentMetadata,omitempty"`
    IsFolder *bool `json:"isFolder,omitempty"`
    IsSymLink *bool `json:"isSymLink,omitempty"`
    Path *string `json:"path,omitempty"`
    Url *string `json:"url,omitempty"`
    ChangeDate *time.Time `json:"changeDate,omitempty"`
    DeletionId *int `json:"deletionId,omitempty"`
    // File encoding from database, -1 represents binary.
    Encoding *int `json:"encoding,omitempty"`
    // MD5 hash as a base 64 string, applies to files only.
    HashValue *string `json:"hashValue,omitempty"`
    IsBranch *bool `json:"isBranch,omitempty"`
    IsPendingChange *bool `json:"isPendingChange,omitempty"`
    // The size of the file, if applicable.
    Size *uint64 `json:"size,omitempty"`
    Version *int `json:"version,omitempty"`
    // MD5 hash as a base 64 string, applies to files only.
    PreviousHashValue *string `json:"previousHashValue,omitempty"`
}

type TfvcItemRequestData struct {
    // If true, include metadata about the file type
    IncludeContentMetadata *bool `json:"includeContentMetadata,omitempty"`
    // Whether to include the _links field on the shallow references
    IncludeLinks *bool `json:"includeLinks,omitempty"`
    ItemDescriptors *[]TfvcItemDescriptor `json:"itemDescriptors,omitempty"`
}

type TfvcLabel struct {
    Links *ReferenceLinks `json:"_links,omitempty"`
    Description *string `json:"description,omitempty"`
    Id *int `json:"id,omitempty"`
    LabelScope *string `json:"labelScope,omitempty"`
    ModifiedDate *time.Time `json:"modifiedDate,omitempty"`
    Name *string `json:"name,omitempty"`
    Owner *IdentityRef `json:"owner,omitempty"`
    Url *string `json:"url,omitempty"`
    Items *[]TfvcItem `json:"items,omitempty"`
}

type TfvcLabelRef struct {
    Links *ReferenceLinks `json:"_links,omitempty"`
    Description *string `json:"description,omitempty"`
    Id *int `json:"id,omitempty"`
    LabelScope *string `json:"labelScope,omitempty"`
    ModifiedDate *time.Time `json:"modifiedDate,omitempty"`
    Name *string `json:"name,omitempty"`
    Owner *IdentityRef `json:"owner,omitempty"`
    Url *string `json:"url,omitempty"`
}

type TfvcLabelRequestData struct {
    // Whether to include the _links field on the shallow references
    IncludeLinks *bool `json:"includeLinks,omitempty"`
    ItemLabelFilter *string `json:"itemLabelFilter,omitempty"`
    LabelScope *string `json:"labelScope,omitempty"`
    MaxItemCount *int `json:"maxItemCount,omitempty"`
    Name *string `json:"name,omitempty"`
    Owner *string `json:"owner,omitempty"`
}

type TfvcMappingFilter struct {
    Exclude *bool `json:"exclude,omitempty"`
    ServerPath *string `json:"serverPath,omitempty"`
}

type TfvcMergeSource struct {
    // Indicates if this a rename source. If false, it is a merge source.
    IsRename *bool `json:"isRename,omitempty"`
    // The server item of the merge source
    ServerItem *string `json:"serverItem,omitempty"`
    // Start of the version range
    VersionFrom *int `json:"versionFrom,omitempty"`
    // End of the version range
    VersionTo *int `json:"versionTo,omitempty"`
}

type TfvcPolicyFailureInfo struct {
    Message *string `json:"message,omitempty"`
    PolicyName *string `json:"policyName,omitempty"`
}

type TfvcPolicyOverrideInfo struct {
    Comment *string `json:"comment,omitempty"`
    PolicyFailures *[]TfvcPolicyFailureInfo `json:"policyFailures,omitempty"`
}

type TfvcShallowBranchRef struct {
    // Path for the branch.
    Path *string `json:"path,omitempty"`
}

// This is the deep shelveset class
type TfvcShelveset struct {
    Links *ReferenceLinks `json:"_links,omitempty"`
    Comment *string `json:"comment,omitempty"`
    CommentTruncated *bool `json:"commentTruncated,omitempty"`
    CreatedDate *time.Time `json:"createdDate,omitempty"`
    Id *string `json:"id,omitempty"`
    Name *string `json:"name,omitempty"`
    Owner *IdentityRef `json:"owner,omitempty"`
    Url *string `json:"url,omitempty"`
    Changes *[]TfvcChange `json:"changes,omitempty"`
    Notes *[]CheckinNote `json:"notes,omitempty"`
    PolicyOverride *TfvcPolicyOverrideInfo `json:"policyOverride,omitempty"`
    WorkItems *[]AssociatedWorkItem `json:"workItems,omitempty"`
}

// This is the shallow shelveset class
type TfvcShelvesetRef struct {
    Links *ReferenceLinks `json:"_links,omitempty"`
    Comment *string `json:"comment,omitempty"`
    CommentTruncated *bool `json:"commentTruncated,omitempty"`
    CreatedDate *time.Time `json:"createdDate,omitempty"`
    Id *string `json:"id,omitempty"`
    Name *string `json:"name,omitempty"`
    Owner *IdentityRef `json:"owner,omitempty"`
    Url *string `json:"url,omitempty"`
}

type TfvcShelvesetRequestData struct {
    // Whether to include policyOverride and notes Only applies when requesting a single deep shelveset
    IncludeDetails *bool `json:"includeDetails,omitempty"`
    // Whether to include the _links field on the shallow references. Does not apply when requesting a single deep shelveset object. Links will always be included in the deep shelveset.
    IncludeLinks *bool `json:"includeLinks,omitempty"`
    // Whether to include workItems
    IncludeWorkItems *bool `json:"includeWorkItems,omitempty"`
    // Max number of changes to include
    MaxChangeCount *int `json:"maxChangeCount,omitempty"`
    // Max length of comment
    MaxCommentLength *int `json:"maxCommentLength,omitempty"`
    // Shelveset's name
    Name *string `json:"name,omitempty"`
    // Owner's ID. Could be a name or a guid.
    Owner *string `json:"owner,omitempty"`
}

type TfvcStatistics struct {
    // Id of the last changeset the stats are based on.
    ChangesetId *int `json:"changesetId,omitempty"`
    // Count of files at the requested scope.
    FileCountTotal *uint64 `json:"fileCountTotal,omitempty"`
}

type TfvcVersionDescriptor struct {
    Version *string `json:"version,omitempty"`
    VersionOption *TfvcVersionOption `json:"versionOption,omitempty"`
    VersionType *TfvcVersionType `json:"versionType,omitempty"`
}

type TfvcVersionOption string

type TfvcVersionType string

// Real time event (SignalR) for a title/description update on a pull request
type TitleDescriptionUpdatedEvent struct {
    // The id of this event. Can be used to track send/receive state between client and server.
    EventId *uuid.UUID `json:"eventId,omitempty"`
    // The id of the pull request this event was generated for.
    PullRequestId *int `json:"pullRequestId,omitempty"`
}

type UpdateRefsRequest struct {
    RefUpdateRequests *[]GitRefUpdate `json:"refUpdateRequests,omitempty"`
    UpdateMode *GitRefUpdateMode `json:"updateMode,omitempty"`
}

type VersionControlChangeType string

type VersionControlProjectInfo struct {
    DefaultSourceControlType *SourceControlTypes `json:"defaultSourceControlType,omitempty"`
    Project *TeamProjectReference `json:"project,omitempty"`
    SupportsGit *bool `json:"supportsGit,omitempty"`
    SupportsTFVC *bool `json:"supportsTFVC,omitempty"`
}

type VersionControlRecursionType string

// A particular revision for a policy configuration.
type VersionedPolicyConfigurationRef struct {
    // The policy configuration ID.
    Id *int `json:"id,omitempty"`
    // The policy configuration type.
    Type *PolicyTypeRef `json:"type,omitempty"`
    // The URL where the policy configuration can be retrieved.
    Url *string `json:"url,omitempty"`
    // The policy configuration revision ID.
    Revision *int `json:"revision,omitempty"`
}

// The representation of data needed to create a tag definition which is sent across the wire.
type WebApiCreateTagRequestData struct {
    // Name of the tag definition that will be created.
    Name *string `json:"name,omitempty"`
}

// The representation of a tag definition which is sent across the wire.
type WebApiTagDefinition struct {
    // Whether or not the tag definition is active.
    Active *bool `json:"active,omitempty"`
    // ID of the tag definition.
    Id *uuid.UUID `json:"id,omitempty"`
    // The name of the tag definition.
    Name *string `json:"name,omitempty"`
    // Resource URL for the Tag Definition.
    Url *string `json:"url,omitempty"`
}
